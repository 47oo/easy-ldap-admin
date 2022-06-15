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
	"log"

	"github.com/spf13/cobra"
)

var teammodDesc string

func teammodRun(cmd *cobra.Command, args []string) {
	o := eldap.NewOption()
	t := eldap.NewTeamEntry()
	t.Name = args
	t.Description = append(t.Description, teammodDesc)
	if err := o.TeamDescUpdate(t); err != nil {
		log.Fatalln(err)
	}
}

// teammodCmd represents the teammod command
var teammodCmd = &cobra.Command{
	Use:   "teammod [flags] TEAM",
	Short: "modify a team",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		teammodRun(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(teammodCmd)
	teammodCmd.Flags().StringVarP(&teammodDesc, "desc", "d", "", "Team Description")
	teammodCmd.MarkFlagRequired("desc")

}
