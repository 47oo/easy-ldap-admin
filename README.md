# Easy LDAP Admin

## How to make

```shell
git clone https://github.com/47oo/easy-ldap-admin.git
cd easy-ldap-admin
make build
```

The binary is output to `bin/ela`.

Other make targets:

| Command          | Description                                            |
|------------------|--------------------------------------------------------|
| `make build`     | Build the binary into `./bin/`                         |
| `make build-all` | Cross-compile linux/darwin/windows (amd64 + arm64)     |
| `make test`      | Run all tests                                          |
| `make vet`       | Run go vet                                             |
| `make fmt`       | Format all Go code                                     |
| `make clean`     | Remove build artifacts (`./bin/`)                      |
| `make help`      | List available targets                                 |

Requirements: Go 1.17+, GNU Make. `make build-all` also needs no extra tooling — it just sets `GOOS`/`GOARCH` per target.

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
  usermod   

Flags:
      --config string   config file (default is $HOME/.ela.ini)
  -h, --help            help for ela
  -t, --toggle          Help message for toggle

Use "ela [command] --help" for more information about a command.
```
