package main

import (
	"os"

	"github.com/Mohammad-y-abbass/react-cli/cmd"
)

// main is the entry point of the React CLI application
func main() {
	// Execute the root command
	cmd.Execute()

	// Exit with success code
	os.Exit(0)
}
