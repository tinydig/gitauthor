package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func Current() {
	name, err := exec.Command("git", "config", "user.name").Output()
	if len(name) != 0 && err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	nameStr := string(name)
	if len(nameStr) == 0 {
		nameStr = "\n"
	}	

	email, err := exec.Command("git", "config", "user.email").Output()
	if len(email) != 0 && err != nil {
		fmt.Println("ss", email, err)
		os.Exit(1)
	}

	emailStr := string(email)
	if len(emailStr) == 0 {
		emailStr = "\n"
	}

	fmt.Printf("name: %v", nameStr)
	fmt.Printf("email: %v", emailStr)
}
