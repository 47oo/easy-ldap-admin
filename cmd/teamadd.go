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

var teamaddName string
var teamaddDesc string

func teamaddRun() {
	o := eldap.NewOption()
	t := eldap.NewTeamEntry()
	t.Name = append(t.Name, teamaddName)
	t.Description = append(t.Description, teamaddDesc)
	if err := o.TeamAdd(t); err != nil {
		log.Fatalln(err)
	}

}

// teamaddCmd represents the teamadd command
var teamaddCmd = &cobra.Command{
	Use:   "teamadd",
	Short: "create a new team",
	Long: `team is an organization in ldap. For example:

ela teamadd -n <teamname> [-d <descprition>]`,
	Run: func(cmd *cobra.Command, args []string) {
		teamaddRun()
	},
}

func init() {
	rootCmd.AddCommand(teamaddCmd)

	teamaddCmd.Flags().StringVarP(&teamaddName, "name", "n", "", "Team Name you create")
	teamaddCmd.Flags().StringVarP(&teamaddDesc, "desc", "d", "no_desc", "Team Description")
	teamaddCmd.MarkFlagRequired("name")

}
