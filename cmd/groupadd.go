/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

var groupaddInfo = model.GroupInfo{}

func groupaddRun() {

	o := eldap.NewOption()
	if err := o.GroupAdd(groupaddInfo); err != nil {
		log.Fatalln(err)
	}

}

// groupaddCmd represents the groupadd command
var groupaddCmd = &cobra.Command{
	Use:   "groupadd",
	Short: "create a group",
	Long: `The groupadd command creates a new group account using the values specified on the command line plus the default values from
	the system. The new group will be entered into the ldap server as needed.`,
	Run: func(cmd *cobra.Command, args []string) {
		groupaddRun()
	},
}

func init() {
	rootCmd.AddCommand(groupaddCmd)
	groupaddCmd.Flags().StringVarP(&groupaddInfo.GidNumber, "gid", "g", "", "use GID for the new group")
	groupaddCmd.Flags().StringVarP(&groupaddInfo.Name, "name", "n", "", "Group Name")
	groupaddCmd.Flags().StringVarP(&groupaddInfo.Description, "desc", "d", "no_desc", "Group Description")
	groupaddCmd.Flags().StringVarP(&groupaddInfo.TeamName, "teamname", "t", "", "You want the group in which team, or default team")
	groupaddCmd.MarkFlagRequired("name")
	groupaddCmd.MarkFlagRequired("gid")

}
