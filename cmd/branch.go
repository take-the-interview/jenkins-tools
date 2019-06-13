// Copyright © 2019 Alen Komic <akomic@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"jenkins-tools/branch"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Branch related",

	Example: `
export BRANCH_NAME=...

jenkins-tools github branch --name "i20" --create-branch-from "staging"
`,

	Run: githubBranchRun,
}

func init() {
	githubCmd.AddCommand(branchCmd)

	branchCmd.Flags().StringP("name", "", "", "Branch name BRANCH_NAME")
	branchCmd.Flags().BoolP("get-deployment-env", "", false, "Get deployment environment name")
	branchCmd.Flags().StringP("create-branch-from", "", "", "Create new branch from")

	viper.BindPFlag("name", branchCmd.Flags().Lookup("name"))
	viper.BindPFlag("get-deployment-env", branchCmd.Flags().Lookup("get-deployment-env"))
	viper.BindPFlag("create-branch-from", branchCmd.Flags().Lookup("create-branch-from"))

	viper.BindEnv("name", "BRANCH_NAME")
}

func branchIsAllowed(branchName string) bool {
	allowedBranches := []string{
		"i1",
		"i2",
		"i3",
		"i4",
		"i5",
		"i6",
		"i7",
		"i8",
		"i9",
		"i10",
		"i11",
		"i12",
		"i13",
		"i14",
		"i15",
		"i16",
		"i17",
		"i18",
		"i19",
		"i20",
		"is1",
		"is2",
		"is3",
		"is4",
		"is5",
	}

	for _, b := range allowedBranches {
		if b == branchName {
			return true
		}
	}
	return false
}

func githubBranchRun(cmd *cobra.Command, args []string) {
	owner := viper.GetString("owner")
	repo := viper.GetString("repo")
	name := viper.GetString("name")

	if owner == "" || repo == "" || name == "" {
		fmt.Println("--name --owner --repo required!")
		os.Exit(1)
	}

	getDeploymentEnv := viper.GetBool("get-deployment-env")
	createBranchFrom := viper.GetString("create-branch-from")

	if getDeploymentEnv {
		branch.GetDeploymentEnv()
	} else if createBranchFrom != "" {
		if !branchIsAllowed(name) {
			fmt.Printf("Branch %s is not allowed.\n", name)
		} else {
			branch.CreateBranch()
		}
	} else {
		cmd.Help()
		os.Exit(0)
	}
}
