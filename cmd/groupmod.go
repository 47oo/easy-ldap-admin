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

var groupmodGidNumber string
var groupmodDesc string

func groupmodRun(cmd *cobra.Command, args []string) {
	o := eldap.NewOption()
	if groupaddGidNumber != "" {
		if err := o.GroupMod(args[0], groupmodGidNumber); err != nil {
			log.Fatalln(err)
		}
	}

}

// groupmodCmd represents the groupmod command
var groupmodCmd = &cobra.Command{
	Use:   "groupmod [flags] GROUP",
	Short: "modify a group definition on the system",
	Long:  `The groupmod command modifies the definition of the specified GROUP by modifying the appropriate entry in the group database.`,
	Run: func(cmd *cobra.Command, args []string) {
		groupmodRun(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(groupmodCmd)
	groupmodCmd.Flags().StringVarP(&groupmodGidNumber, "gid", "g", "", "change the group ID to GID")
	groupmodCmd.Flags().StringVarP(&groupmodDesc, "desc", "d", "", "Descroption Not support now")

}
