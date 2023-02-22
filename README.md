# Gitauthor

Quickly setup git repository author details.

## Install & run

Requires Golang for building packages!

1. Clone the repository
2. Run `sh ./scripts/local.sh`

## Commands
| Command | Details | Example |
| --- | --- | --- |
| help | Print help menu | `gitauthor help` |
| init | Creates config file | `gitauthor init` | 
| config | List path to config file | `gitauthor config` |
| list | List all configured authors from config file | `gitauthor list` |
| select | Apply an author to current repository | `gitauthor select [user]` |
| current | List current git author for current repo | `gitauthor current` |

## License

- [MIT License](LICENSE)