package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/tinydig/gitauthor/internal/config"
)

func Select() {
	argsLen := len(os.Args[1:])
	if argsLen < 2 {
		fmt.Println("user not specified")
		os.Exit(1)
	}

	users := config.GetAuthorsList()
	user := users[os.Args[2]]

	if (user == config.Author{}) {
		fmt.Printf("user %v not found\n", os.Args[2])
		os.Exit(1)
	}

	formattedValue := fmt.Sprintf("\"%v\"", user.Name)
	err := exec.Command("git", "config", "user.name", formattedValue).Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	formattedValue = fmt.Sprintf("\"%v\"", user.Email)
	err = exec.Command("git", "config", "user.email", formattedValue).Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
