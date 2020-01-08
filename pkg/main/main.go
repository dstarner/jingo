package main

import (
	"github.com/dstarner/jingo/cmd"
)

// The main entrance-point of the CLI application
func main() {
	rootCommand := cmd.NewJingoCommand()
	if err := rootCommand.Execute(); err != nil {
		
	}
}
