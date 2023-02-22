# Gitauthor

Quickly setup git repository author details.

## Install

### Install for Mac

```
curl -sSfL https://raw.githubusercontent.com/tinydig/gitauthor/master/scripts/install.sh | sh -s v0.1.0-alpha.2 darwin

gitauthor init
```

### Manually install for Mac

Download the binary from release and put it in `/usr/local/bin`.

### Install for Linux

```
curl -sSfL https://raw.githubusercontent.com/tinydig/gitauthor/master/scripts/install.sh | sh -s v0.1.0-alpha.2 linux

gitauthor init
```

### Manually install for Linux

Download the binary from release and put it in `/usr/local/bin`.

### Install for Windows

Download the binary and put it somewhere in your %path%.

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