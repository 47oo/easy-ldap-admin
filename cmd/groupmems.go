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

var groupname string
var addusername string
var delusername string
var groupmemslist bool

func groupmemsRun() {
	o := eldap.NewOption()
	if groupmemslist {
		log.Fatalln("not support now")
		return
	}
	if addusername != "" {
		if err := o.GroupMems(groupname, []string{addusername}, eldap.Add); err != nil {
			log.Fatalln(err)
			return
		}
	}
	if delusername != "" {
		if err := o.GroupMems(groupname, []string{delusername}, eldap.Del); err != nil {
			log.Fatalln(err)
			return
		}
	}

}

// groupmemsCmd represents the groupmems command
var groupmemsCmd = &cobra.Command{
	Use:   "groupmems",
	Short: "administer members of a user's primary group",
	Long: `The groupmems command allows a user to administer their own group membership list without the requirement of superuser
	privileges`,
	Run: func(cmd *cobra.Command, args []string) {
		groupmemsRun()
	},
}

func init() {
	rootCmd.AddCommand(groupmemsCmd)
	groupmemsCmd.Flags().StringVarP(&groupname, "group", "g", "", "change groupname instead of the user's group")
	groupmemsCmd.Flags().StringVarP(&addusername, "add", "a", "", "add username to the members of the group")
	groupmemsCmd.Flags().StringVarP(&delusername, "delete", "d", "", "add username to the members of the group")
	groupmemsCmd.Flags().BoolVarP(&groupmemslist, "list", "l", false, "list the members of the group")

	groupmemsCmd.MarkFlagRequired("group")

}
