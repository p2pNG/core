package main

import (
	"fmt"
	"github.com/p2pNG/core/cmd/commands"
)

func main() {
	err := commands.Execute()
	if err != nil {
		fmt.Printf("%v", err)
	}
}
