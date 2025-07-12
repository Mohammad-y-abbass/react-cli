package cmd

import (
	"os"

	"github.com/Mohammad-y-abbass/react-cli/helpers"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rc",
	Short: "React CLI - A modern CLI tool for creating React projects",
	Long: `React CLI is a command-line tool that simplifies the process of creating 
React applications with modern tooling.

Features:
  • Create React apps with Vite
  • TypeScript support
  • Tailwind CSS integration
  • Modern project structure

Examples:
  rc new my-app              # Create a basic React app
  rc new my-app --ts         # Create a React app with TypeScript
  rc new my-app --tailwind   # Create a React app with Tailwind CSS`,
	Version: "1.0.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		helpers.GlobalLogger.Error("Error: " + err.Error())
		os.Exit(1)
	}
}

func init() {
	// Add any global flags here if needed in the future
}
