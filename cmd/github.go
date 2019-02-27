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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Github integrations",
	Long:  `Various things to help integrate Jenkins and Github`,
}

func init() {
	rootCmd.AddCommand(githubCmd)

	githubCmd.PersistentFlags().StringP("auth-token", "", "", "Github auth token GITHUB_AUTH_TOKEN")
	githubCmd.PersistentFlags().StringP("owner", "", "", "Github owner GITHUB_OWNER")
	githubCmd.PersistentFlags().StringP("repo", "", "", "Github owner GITHUB_REPO")

	viper.BindPFlag("auth-token", githubCmd.PersistentFlags().Lookup("auth-token"))
	viper.BindPFlag("owner", githubCmd.PersistentFlags().Lookup("owner"))
	viper.BindPFlag("repo", githubCmd.PersistentFlags().Lookup("repo"))

	viper.BindEnv("auth-token", "GITHUB_AUTH_TOKEN")
	viper.BindEnv("owner", "GITHUB_OWNER")
	viper.BindEnv("repo", "GITHUB_REPO")
}
