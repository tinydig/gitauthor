package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/tinydig/gitauthor/internal/config"
)

func folderExists(folder string) (bool, error) {
	_, err := os.Stat(folder)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func initConfigDir() {
	configDir := config.ConfigDir()
	exists, err := folderExists(configDir)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if exists != true {
		os.Mkdir(configDir, os.FileMode(int(0755)))
	}
}

func initConfigFile() {
	configFile := config.ConfigPath()
	exists, err := folderExists(configFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if exists != true {
		srcFile, err1 := os.Open("./assets/users.yaml")
		if err1 != nil {
			fmt.Println(err1)
			os.Exit(1)
		}
		defer srcFile.Close()

		destFile, err2 := os.Create(configFile)
		if err2 != nil {
			fmt.Println(err2)
			os.Exit(1)
		}
		defer destFile.Close()

		_, err3 := io.Copy(destFile, srcFile)
		err3 = destFile.Sync()
		if err3 != nil {
			fmt.Println(err3)
			os.Exit(1)
		}
	}
}

func Init() {
	initConfigDir()
	initConfigFile()
}
