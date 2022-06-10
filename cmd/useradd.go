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

var useraddInfo = model.UserInfo{}

func useraddRun() {
	o := eldap.NewOption()
	if err := o.UserAdd(useraddInfo); err != nil {
		log.Fatalln(err)
	}
}

// useraddCmd represents the useradd command
var useraddCmd = &cobra.Command{
	Use:   "useradd",
	Short: "create a new user or update default new user information",
	Long: `the useradd command creates a new user account using the values specified on the command
	line plus the default values from the ldap`,
	Run: func(cmd *cobra.Command, args []string) {
		useraddRun()
	},
}

func init() {
	rootCmd.AddCommand(useraddCmd)
	useraddCmd.Flags().StringVarP(&useraddInfo.Name, "name", "n", "", "username")
	useraddCmd.Flags().StringVarP(&useraddInfo.HomeDirectory, "home-dir", "d", "", "user home dir")
	useraddCmd.Flags().StringVarP(&useraddInfo.GidNumber, "gid", "g", "", "group number")
	useraddCmd.Flags().StringVarP(&useraddInfo.UidNumber, "uid", "u", "", "user uid number")
	useraddCmd.Flags().StringVarP(&useraddInfo.UserPassword, "password", "p", "", "encrypted password of the new account")
	useraddCmd.Flags().StringVarP(&useraddInfo.LoginShell, "shell", "s", "", "login shell of the new account")
	useraddCmd.Flags().StringVarP(&useraddInfo.TeamName, "team", "t", "", "teamname for this user")

	useraddCmd.MarkFlagRequired("name")
	useraddCmd.MarkFlagRequired("home-dir")
	useraddCmd.MarkFlagRequired("gid")
	useraddCmd.MarkFlagRequired("uid")
	useraddCmd.MarkFlagRequired("password")
	useraddCmd.MarkFlagRequired("shell")
}
