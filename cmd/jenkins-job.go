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

	"jenkins-tools/jenkins"
)

var jenkinsjobCmd = &cobra.Command{
	Use:   "job",
	Short: "Job related",

	Example: `jenkins-tools jenkins job --name "Back-End/job/sms-service/job/i9" --get-commit-status "595be0f62eb9ef984cd8f4e63d6956b84131a6d0"`,

	Run: jenkinsJobRun,
}

func init() {
	jenkinsCmd.AddCommand(jenkinsjobCmd)

	jenkinsjobCmd.Flags().StringP("name", "", "", "Job Name")
	jenkinsjobCmd.Flags().StringP("get-commit-status", "", "", "Get status of the job for the given Git Commit SHA (with waiting)")

	viper.BindPFlag("name", jenkinsjobCmd.Flags().Lookup("name"))
	viper.BindPFlag("get-commit-status", jenkinsjobCmd.Flags().Lookup("get-commit-status"))
}

func jenkinsJobRun(cmd *cobra.Command, args []string) {
	name := viper.GetString("name")
	getCommitStatus := viper.GetString("get-commit-status")

	if name == "" {
		fmt.Println("--name required!")
		os.Exit(1)
	}

	if getCommitStatus != "" {
		jenkins.CommitStatus()
	} else {
		cmd.Help()
		os.Exit(0)
	}

	// getDeploymentEnv := viper.GetBool("get-deployment-env")
	// createBranchFrom := viper.GetString("create-branch-from")

	// if getDeploymentEnv {
	// 	branch.GetDeploymentEnv()
	// } else if createBranchFrom != "" {
	// 	if !branchIsAllowed(name) {
	// 		fmt.Printf("Branch %s is not allowed.\n", name)
	// 	} else {
	// 		branch.CreateBranch()
	// 	}
	// } else {
	// 	cmd.Help()
	// 	os.Exit(0)
	// }
}
