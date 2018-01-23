package main

import (
	"fmt"
	"os"

	"github.com/somyagarg94/error_calculator/cmd/cli/command"
)

// Set via linker flag
var version string
var buildDate string

func main() {
	error_calculatorCliCmd, err := command.NewError_calculatorCliCommand(version, buildDate, os.Stdin, os.Stdout)
	if err != nil {
		fmt.Printf("Error initializing command: %v\n", err)
		os.Exit(1)
	}
	if err = error_calculatorCliCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
