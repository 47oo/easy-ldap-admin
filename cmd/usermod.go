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

var usermodHome string
var usermodName string
var usermodPassword string
var usermodShell string
var usermodUidNumber string
var usermodGidNumber string

func usermodRun() {
	o := eldap.NewOption()
	u := eldap.NewUserEntry()
	u.Name = append(u.Name, usermodName)
	u.LoginShell = append(u.LoginShell, usermodShell)
	u.GidNumber = append(u.GidNumber, usermodGidNumber)
	u.HomeDirectory = append(u.HomeDirectory, usermodHome)
	u.UserPassword = append(u.UserPassword, usermodPassword)
	u.UidNumber = append(u.UidNumber, usermodUidNumber)
	if err := o.UserMod(u); err != nil {
		log.Fatalln(err)
	}

}

// usermodCmd represents the usermod command
var usermodCmd = &cobra.Command{
	Use:   "usermod",
	Short: "modify a user account",
	Long:  `he usermod command modifies the ldap to reflect the changes that are specified on the command line`,
	Run: func(cmd *cobra.Command, args []string) {
		usermodRun()
	},
}

func init() {
	rootCmd.AddCommand(usermodCmd)
	usermodCmd.Flags().StringVarP(&usermodHome, "home", "d", "", "new home directory for the user account")
	usermodCmd.Flags().StringVarP(&usermodName, "name", "n", "", "which account you want to change")
	usermodCmd.Flags().StringVarP(&usermodPassword, "password", "p", "", "use encrypted password for the new password")
	usermodCmd.Flags().StringVarP(&usermodShell, "shell", "s", "", "new login shell for the user account")
	usermodCmd.Flags().StringVarP(&usermodUidNumber, "uid", "u", "", "new UID for the user account")
	usermodCmd.Flags().StringVarP(&usermodGidNumber, "gid", "g", "", "force use GID as new primary group")
	usermodCmd.MarkFlagRequired("name")
}
