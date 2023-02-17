package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Author struct {
	Name  string
	Email string
}

func GetAuthorsList() map[string]Author {
	configPath := ConfigPath()
	yfile, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Fatal(err)
	}

	authorsDetails := make(map[string]Author)
	err2 := yaml.Unmarshal(yfile, &authorsDetails)

	if err2 != nil {
		log.Fatal(err2)
	}

	return authorsDetails
}
