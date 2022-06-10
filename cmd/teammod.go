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

func teammodRun() {
	o := eldap.NewOption()
	if teammodInfo.Name == "" || teammodInfo.Description == "" {
		log.Fatalln("TeamName Must exist or Desc Must exist")
		return
	}
	if err := o.TeamDescUpdate(teammodInfo); err != nil {
		log.Fatalln(err)
	}
}

// teammodCmd represents the teammod command
var teammodCmd = &cobra.Command{
	Use:   "teammod",
	Short: "modify a team",
	Long:  `modify a team,only support modify desc`,
	Run: func(cmd *cobra.Command, args []string) {
		teammodRun()
	},
}

var teammodInfo = model.TeamInfo{}

func init() {
	rootCmd.AddCommand(teammodCmd)
	teammodCmd.Flags().StringVarP(&teammodInfo.Name, "name", "n", "", "Team Name")
	teammodCmd.Flags().StringVarP(&teammodInfo.Description, "desc", "d", "", "Team Description")
	teammodCmd.MarkFlagRequired("name")
	teammodCmd.MarkFlagRequired("desc")

}
