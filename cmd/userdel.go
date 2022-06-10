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

var userdelInfo = model.UserInfo{}

func userdelRun() {
	o := eldap.NewOption()
	if err := o.UserDel(userdelInfo.Name); err != nil {
		log.Fatalln(err)
	}
}

// userdelCmd represents the userdel command
var userdelCmd = &cobra.Command{
	Use:   "userdel",
	Short: "delete a user account and related files",
	Long:  `The userdel command modifies ldap server data,  The named user must exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		userdelRun()
	},
}

func init() {
	rootCmd.AddCommand(userdelCmd)
	userdelCmd.Flags().StringVarP(&userdelInfo.Name, "name", "n", "", "username you want to delete")
	userdelCmd.MarkFlagRequired("name")

}
