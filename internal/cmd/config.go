package cmd

import (
	"fmt"

	"github.com/tinydig/gitauthor/internal/config"
)

func Config() {
	fmt.Println(config.ConfigPath())
}
