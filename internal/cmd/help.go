package cmd

import "fmt"

func Help() {
	fmt.Println("Commands:")
	fmt.Println("  help: print this message :)")
	fmt.Println("  init: init config file if not exists")
	fmt.Println("  current: list current author configuration")
	fmt.Println("  list: list configured authors")
	fmt.Println("  config: print authors config file path")
	fmt.Println("  select: set specified author to current git repository")
}
