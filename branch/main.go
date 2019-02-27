package branch

import (
	"fmt"
	"os"

	"jenkins-tools/config"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"strings"
)

// GetDeploymentEnv Trying to figure out if the specified owner/repo/branch
// has pull request in open state with label "deploy:<env name>"
// Return empty string otherwise, even in case of error
func GetDeploymentEnv() {
	authToken := viper.GetString("auth-token")
	owner := viper.GetString("owner")
	repo := viper.GetString("repo")
	branchName := viper.GetString("name")

	config.Auth(authToken)

	opts := &github.PullRequestListOptions{
		State: "open",
		Head:  fmt.Sprintf("%s:%s", owner, branchName),
	}
	prs, _, err := config.Client.PullRequests.List(config.Ctx, owner, repo, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		return
	}

	for _, pr := range prs {
		for _, label := range pr.Labels {
			labelChunks := strings.Split(*label.Name, ":")
			if len(labelChunks) == 2 && labelChunks[0] == "deploy" {
				fmt.Printf("%s\n", labelChunks[1])
				return
			}
		}
	}
}
