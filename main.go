package main

import (
	"github.com/chrisgavin/gh-sops/cmd"
)

func main() {
	rootCommand, err := cmd.NewRootCommand()
	if err != nil {
		panic(err)
	}
	rootCommand.Run()
}
