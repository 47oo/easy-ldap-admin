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
	"ela/model"
	"ela/secret"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

func initRun() {
	lai := model.LDAPAuthInfo{}
	fmt.Printf("Please enter your ldap server host: ")
	fmt.Scanln(&lai.LDAPHost)
	fmt.Printf("Please enter user ldap server port: ")
	fmt.Scanln(&lai.LDAPPort)
	fmt.Printf("Please enter ldap Top DN: ")
	fmt.Scanln(&lai.TopDN)
	fmt.Printf("Please enter ldap Admin account: ")
	fmt.Scanln(&lai.Admin)
	pd, _ := gopass.GetPasswdPrompt(`Please enter ldap admin passwd(enter "NO" to not write password): `, true, os.Stdin, os.Stdout)
	asepd, err := secret.EncryptAES([]byte(pd), secret.KEY)
	if err != nil {
		log.Fatalln(err)
		return
	}

	lai.AdminPW = string(asepd)
	homedir, _ := os.UserHomeDir()
	filename := path.Join(homedir, ".ela.ini")
	cfg := ini.Empty()
	dS := cfg.Section("")
	dS.NewKey("LDAPHost", lai.LDAPHost)
	dS.NewKey("LDAPPort", lai.LDAPPort)
	dS.NewKey("Admin", lai.Admin)
	dS.NewKey("AdminPW", lai.AdminPW)
	dS.NewKey("TopDN", lai.TopDN)
	if err = cfg.SaveTo(filename); err != nil {
		log.Fatalln(err)
	}

}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init ldap config",
	Long: `init config with connect to ldap,the config save at $HOME/.ela.ini For example:
ela init`,
	Run: func(cmd *cobra.Command, args []string) {
		initRun()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}
