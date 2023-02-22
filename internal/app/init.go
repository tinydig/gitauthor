package app

import (
	"fmt"
	"os"

	"github.com/tinydig/gitauthor/internal/cmd"
)

func Start() {
	argsLen := len(os.Args[1:])
	if argsLen < 1 {
		cmd.Help()
		return
	}

	command := os.Args[1]
	if command == "help" {
		cmd.Help()
		return
	}

	if command == "init" {
		cmd.Init()
		return
	}

	if command == "list" {
		cmd.List()
		return
	}

	if command == "select" {
		cmd.Select()
		cmd.Current()
		return
	}

	if command == "config" {
		cmd.Config()
		return
	}

	if command == "current" {
		cmd.Current()
		return
	}

	fmt.Println("command not found")
	os.Exit(1)
}
