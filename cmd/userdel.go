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

func userdelRun(cmd *cobra.Command, args []string) {
	o := eldap.NewOption()
	if err := o.UserDel(args[0]); err != nil {
		log.Fatalln(err)
	}
}

// userdelCmd represents the userdel command
var userdelCmd = &cobra.Command{
	Use:   "userdel [flags] LOGIN",
	Short: "delete a user account and related files",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userdelRun(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(userdelCmd)
}
