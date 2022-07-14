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
	"strconv"

	"github.com/spf13/cobra"
)

var groupaddGidNumber string
var groupaddDesc string
var groupaddTeamName string

func groupaddRun(cmd *cobra.Command, args []string) {

	o := eldap.NewOption()
	g := eldap.NewGroupEntry()
	g.Name = args
	g.Description = append(g.Description, groupaddDesc)
	if groupaddGidNumber == "" {
		g, err := eldap.NewGidNumber(eldap.MinNumber, eldap.MaxNumber)
		if err != nil {
			log.Fatalln(err)
			return
		}
		groupaddGidNumber = strconv.Itoa(g)
	}
	g.GidNumber = append(g.GidNumber, groupaddGidNumber)
	if err := o.GroupAdd(groupaddTeamName, g); err != nil {
		log.Fatalln(err)
	}

}

// groupaddCmd represents the groupadd command
var groupaddCmd = &cobra.Command{
	Use:   "groupadd [flags] GROUP",
	Short: "create a group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		groupaddRun(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(groupaddCmd)
	groupaddCmd.Flags().StringVarP(&groupaddGidNumber, "gid", "g", "", "use GID for the new group")
	groupaddCmd.Flags().StringVarP(&groupaddDesc, "desc", "d", "", "Group Description")
	groupaddCmd.Flags().StringVarP(&groupaddTeamName, "teamname", "t", "", "You want the group in which team, or default team")

}
