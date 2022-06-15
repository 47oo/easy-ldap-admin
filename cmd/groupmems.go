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

var groupmemsName string
var addUserName string
var delUserName string

func groupmemsRun() {
	o := eldap.NewOption()

	if addUserName != "" {
		if err := o.GroupMems(groupmemsName, []string{addUserName}, eldap.Add); err != nil {
			log.Fatalln(err)
			return
		}
	}
	if delUserName != "" {
		if err := o.GroupMems(groupmemsName, []string{delUserName}, eldap.Del); err != nil {
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
	groupmemsCmd.Flags().StringVarP(&groupmemsName, "group", "g", "", "change groupname instead of the user's group")
	groupmemsCmd.Flags().StringVarP(&addUserName, "add", "a", "", "add username to the members of the group")
	groupmemsCmd.Flags().StringVarP(&delUserName, "delete", "d", "", "add username to the members of the group")
	groupmemsCmd.MarkFlagRequired("group")
}
