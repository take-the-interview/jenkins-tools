package jenkins

import (
	"fmt"
	"os"

	"jenkins-tools/config"

	"github.com/bndr/gojenkins"
	"github.com/spf13/viper"
	"strings"
	"time"
)

var (
	defaultSleepTime  = 20
	defaultRetryCount = 60
)

func tailBuildBySHA(name, SHA string) (err error) {
	var lastBuildNumber int64
	var build *gojenkins.Build
	mySHA := ""

	c := 0
	for {
		if c >= defaultRetryCount {
			err = fmt.Errorf("Unable to find build %s SHA %s", name, SHA)
			return
		}

		c = c + 1

		job, err := config.Jenkins.GetJob(name)
		if err != nil {
			fmt.Printf("Can't find the job %s: %s\n", name, err.Error())
			time.Sleep(time.Duration(defaultSleepTime) * time.Second)
			continue
		}

		build, err = job.GetLastBuild()
		if err != nil {
			fmt.Printf("Can't find the last build for the job %s: %s\n", name, err.Error())
			time.Sleep(time.Duration(defaultSleepTime) * time.Second)
			continue
		}

		buildInfo := build.Info()

		for _, action := range buildInfo.Actions {
			for _, b := range action.LastBuiltRevision.Branch {
				if strings.HasPrefix(b.Name, "refs/remotes/origin") {
					mySHA = b.SHA1
					break
				}
			}
			if mySHA != "" {
				break
			}
		}
		if mySHA == "" {
			fmt.Printf("Can't find the SHA for the job %s: %s\n", name, err.Error())
			time.Sleep(time.Duration(defaultSleepTime) * time.Second)
			continue
		} else if mySHA != SHA {
			fmt.Printf("Wanted SHA %s got %s\n", SHA, mySHA)
			time.Sleep(time.Duration(defaultSleepTime) * time.Second)
			continue
		} else {
			lastBuildNumber = build.GetBuildNumber()
			break
		}
	}

	fmt.Printf("%d SHA is %s\n", lastBuildNumber, mySHA)

	for {
		if build.IsRunning() {
			fmt.Printf("Build %d is running ...\n", lastBuildNumber)
			time.Sleep(time.Duration(defaultSleepTime) * time.Second)
		} else {
			result := build.GetResult()
			fmt.Println(result)
			if result != "SUCCESS" {
				os.Exit(10)
			} else {
				os.Exit(0)
			}
			break
		}
	}
	return
}

func CommitStatus() {
	name := viper.GetString("name")
	serverUrl := viper.GetString("server-url")
	user := viper.GetString("user")
	pass := viper.GetString("pass")

	SHA := viper.GetString("get-commit-status")

	config.JenkinsAuth(serverUrl, user, pass)

	fmt.Printf("Job info for %s\n", name)

	err := tailBuildBySHA(name, SHA)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		return
	}
}
