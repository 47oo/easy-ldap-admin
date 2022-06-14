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
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var useraddName string
var useraddHomeDirectory string
var useraddGidNumber string
var useraddUidNumber string
var useraddUserPassword string
var useraddLoginShell string
var useraddTeamName string

func useraddRun() {
	if useraddHomeDirectory == "" {
		useraddHomeDirectory = fmt.Sprintf(`/%s/%s`, "home", useraddName)
	}
	if useraddUserPassword == "" {
		useraddUserPassword = useraddName
	}
	if useraddLoginShell == "" {
		useraddLoginShell = "/bin/bash"
	}
	o := eldap.NewOption()
	u := eldap.NewUserEntry()
	u.CN = append(u.CN, useraddName)
	u.Name = append(u.Name, useraddName)
	u.GidNumber = append(u.GidNumber, useraddGidNumber)
	u.UidNumber = append(u.UidNumber, useraddUidNumber)
	u.HomeDirectory = append(u.HomeDirectory, useraddHomeDirectory)
	u.LoginShell = append(u.LoginShell, useraddLoginShell)
	u.UserPassword = append(u.UserPassword, useraddUserPassword)

	if err := o.UserAdd(useraddTeamName, u); err != nil {
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
	useraddCmd.Flags().StringVarP(&useraddName, "name", "n", "", "username")
	useraddCmd.Flags().StringVarP(&useraddHomeDirectory, "home-dir", "d", "", "user home dir")
	useraddCmd.Flags().StringVarP(&useraddGidNumber, "gid", "g", "", "group number")
	useraddCmd.Flags().StringVarP(&useraddUidNumber, "uid", "u", "", "user uid number")
	useraddCmd.Flags().StringVarP(&useraddUserPassword, "password", "p", "", "encrypted password of the new account")
	useraddCmd.Flags().StringVarP(&useraddLoginShell, "shell", "s", "", "login shell of the new account")
	useraddCmd.Flags().StringVarP(&useraddTeamName, "team", "t", "", "teamname for this user")

	useraddCmd.MarkFlagRequired("name")
	useraddCmd.MarkFlagRequired("gid")
	useraddCmd.MarkFlagRequired("uid")

}
