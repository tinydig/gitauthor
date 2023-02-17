package cmd

import (
	"fmt"

	"github.com/tinydig/gitauthor/internal/config"
)

func List() {
	users := config.GetAuthorsList()

	for k, v := range users {
		fmt.Printf("%v:\n", k)
		fmt.Printf("  name: %v\n", v.Name)
		fmt.Printf("  email: %v\n", v.Email)
	}
}
