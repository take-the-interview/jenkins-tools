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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// jenkinsCmd represents the jenkins command
var jenkinsCmd = &cobra.Command{
	Use:   "jenkins",
	Short: "jenkins integrations",
	Long:  `Various things to help integrate Jenkins and Jenkins`,

	Example: `
# Jenkins config
export JENKINS_SERVER_URL=http://ci:8080/
export JENKINS_USERNAME=me
export JENKINS_PASSWORD=mypass

jenkins-tools jenkins ...
`,
}

func init() {
	rootCmd.AddCommand(jenkinsCmd)

	jenkinsCmd.PersistentFlags().StringP("server-url", "", "http://localhost:8080", "Jenkins Server URL")
	jenkinsCmd.PersistentFlags().StringP("user", "", "", "Jenkins username")
	jenkinsCmd.PersistentFlags().StringP("pass", "", "", "Jenkins password")

	viper.BindPFlag("server-url", jenkinsCmd.PersistentFlags().Lookup("server-url"))
	viper.BindPFlag("user", jenkinsCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("pass", jenkinsCmd.PersistentFlags().Lookup("pass"))

	viper.BindEnv("server-url", "JENKINS_SERVER_URL")
	viper.BindEnv("user", "JENKINS_USERNAME")
	viper.BindEnv("pass", "JENKINS_PASSWORD")
}
