/*
Copyright © 2022 Hao Han <136698493@qq.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"ela/eldap"
	"ela/model"
	"log"

	"github.com/spf13/cobra"
)

func teamdelRun() {
	o := eldap.NewOption()
	if teamdelInfo.Name == "" {
		log.Fatalln("Team Name Must Exist")
	}
	if err := o.TeamDelete(teamdelInfo.Name); err != nil {
		log.Fatalln(err)
	}
}

// teamdelCmd represents the teamdel command
var teamdelCmd = &cobra.Command{
	Use:   "teamdel",
	Short: "delete a user account",
	Long:  `delete a user account`,
	Run: func(cmd *cobra.Command, args []string) {
		teamdelRun()
	},
}
var teamdelInfo = model.TeamInfo{}

func init() {
	rootCmd.AddCommand(teamdelCmd)
	teamdelCmd.Flags().StringVarP(&teamdelInfo.Name, "name", "n", "", "The team you want to del")
	teamdelCmd.MarkFlagRequired("name")

}
