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

func groupdelRun(cmd *cobra.Command, args []string) {
	o := eldap.NewOption()
	if err := o.GroupDel(args[0]); err != nil {
		log.Fatalln(err)
	}
}

// groupdelCmd represents the groupdel command
var groupdelCmd = &cobra.Command{
	Use:   "groupdel [flags] GROUP",
	Short: "groupdel - delete a group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		groupdelRun(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(groupdelCmd)
}
