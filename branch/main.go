package branch

import (
	"fmt"
	"os"

	"jenkins-tools/config"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	// "reflect"
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

func CreateBranch() {
	authToken := viper.GetString("auth-token")
	owner := viper.GetString("owner")
	repo := viper.GetString("repo")
	sourceBranchName := viper.GetString("create-branch-from")
	branchName := viper.GetString("name")

	config.Auth(authToken)

	refs, _, err := config.Client.Git.GetRefs(config.Ctx, owner, repo, "refs/heads")
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetRefs ERROR: %s\n", err)
		return
	}

	var branches = make(map[string]*github.Reference)

	for _, gitRef := range refs {
		branches[*gitRef.Ref] = gitRef
	}

	sourceRefName := fmt.Sprintf("refs/heads/%s", sourceBranchName)

	if _, ok := branches[sourceRefName]; !ok {
		fmt.Fprintf(os.Stderr, "Source branch %s does not exist.", sourceRefName)
		return
	}

	ref, _, err := config.Client.Git.GetRef(config.Ctx, owner, repo, sourceRefName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetRef ERROR: %s\n", err)
		return
	}
	// fmt.Println("REF: %v", ref)
	// fmt.Println("RES: %v", res)

	newRefName := fmt.Sprintf("refs/heads/%s", branchName)
	newRef := github.Reference{
		Ref:    &newRefName,
		Object: ref.Object,
	}

	if _, ok := branches[newRefName]; !ok {
		ref, _, err = config.Client.Git.CreateRef(config.Ctx, owner, repo, &newRef)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CreateRef ERROR: %s\n", err)
			return
		}
		fmt.Printf("Branch %s created.\n", branchName)
	} else {
		ref, _, err = config.Client.Git.UpdateRef(config.Ctx, owner, repo, &newRef, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "UpdateRef ERROR: %s\n", err)
			return
		}
		fmt.Printf("Branch %s updated.\n", branchName)
	}

	// fmt.Println("REF: %v", ref)
	// fmt.Println("RES: %v", res)
}
