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

var teamaddDesc string

func teamaddRun(cmd *cobra.Command, args []string) {

	o := eldap.NewOption()
	t := eldap.NewTeamEntry()
	t.Name = args
	t.Description = append(t.Description, teamaddDesc)
	if err := o.TeamAdd(t); err != nil {
		log.Fatalln(err)
	}

}

// teamaddCmd represents the teamadd command
var teamaddCmd = &cobra.Command{
	Use:   "teamadd [flags] TEAM",
	Short: "create a new team",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		teamaddRun(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(teamaddCmd)
	teamaddCmd.Flags().StringVarP(&teamaddDesc, "desc", "d", "", "Team Description")

}
