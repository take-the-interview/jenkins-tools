// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"jenkins-tools/branch"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Branch related",

	Run: githubBranchRun,
}

func init() {
	githubCmd.AddCommand(branchCmd)

	branchCmd.Flags().StringP("name", "", "", "Branch name BRANCH_NAME")
	branchCmd.Flags().BoolP("get-deployment-env", "", false, "Get deployment environment name")

	viper.BindPFlag("name", branchCmd.Flags().Lookup("name"))
	viper.BindPFlag("get-deployment-env", branchCmd.Flags().Lookup("get-deployment-env"))

	viper.BindEnv("name", "BRANCH_NAME")
}

func githubBranchRun(cmd *cobra.Command, args []string) {
	owner := viper.GetString("owner")
	repo := viper.GetString("repo")
	name := viper.GetString("name")
	authToken := viper.GetString("auth-token")

	getDeploymentEnv := viper.GetBool("get-deployment-env")

	if getDeploymentEnv {
		branch.GetDeploymentEnv()
	} else {
		fmt.Printf("branch called %s/%s/%s %s\n", owner, repo, name, authToken)
	}
}
