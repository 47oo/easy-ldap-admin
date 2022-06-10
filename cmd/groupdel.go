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

var groupdelInfo = model.GroupInfo{}

func groupdelRun() {
	o := eldap.NewOption()
	if err := o.GroupDel(groupdelInfo.Name); err != nil {
		log.Fatalln(err)
	}
}

// groupdelCmd represents the groupdel command
var groupdelCmd = &cobra.Command{
	Use:   "groupdel",
	Short: "groupdel - delete a group",
	Long:  `The groupdel command modifies the system account files, deleting all entries that refer to GROUP. The named group must exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		groupdelRun()
	},
}

func init() {
	rootCmd.AddCommand(groupdelCmd)
	groupdelCmd.Flags().StringVarP(&groupdelInfo.Name, "name", "n", "", "Group Name")
	groupdelCmd.MarkFlagRequired("name")

}
