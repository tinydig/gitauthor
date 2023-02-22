package config

import (
	"fmt"
	"os"
	"path"
)

const GIT_USER_CONFIG_FOLDER = ".gitauthor"
const GIT_USER_CONFIG_YAML = "users.yaml"

func HomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("home directory could not be found")
		os.Exit(1)
	}

	return dirname
}

func ConfigDir() string {
	return path.Join(HomeDir(), GIT_USER_CONFIG_FOLDER)
}

func ConfigPath() string {
	return path.Join(ConfigDir(), GIT_USER_CONFIG_YAML)
}
