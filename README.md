![easy-ldap-admin](https://socialify.git.ci/47oo/easy-ldap-admin/image?description=1&font=KoHo&language=1&name=1&pattern=Floating%20Cogs&theme=Light)

## How to make

```shell
git pull https://github.com/47oo/easy-ldap-admin.git
cd easy-ldap-admin
go build
```

## How to use

```shell
ela --help
this is a easy cmd for admin to used. For example:

ela init
ela useradd like linux useradd
ela userdel like linux userdel
...

Usage:
  ela [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  groupadd    create a group
  groupdel    groupdel - delete a group
  groupmems   administer members of a user's primary group
  groupmod    A brief description of your command
  help        Help about any command
  init        init ldap config
  teamadd     create a new team
  teamdel     delete a user account
  teammod     modify a team
  useradd     create a new user or update default new user information
  userdel     delete a user account and related files
  usermod     modify user account

Flags:
      --config string   config file (default is $HOME/.ela.ini)
  -h, --help            help for ela
  -t, --toggle          Help message for toggle

Use "ela [command] --help" for more information about a command.
```
