package cmd

import (
	"fmt"
	"strings"

	"github.com/Mohammad-y-abbass/react-cli/helpers"
	"github.com/Mohammad-y-abbass/react-cli/types"
	"github.com/spf13/cobra"
)

// Command flags
var (
	useTypeScript bool
	useTailwind   bool
)

// newCmd represents the new command for creating React applications
var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new React application",
	Long: `Create a new React application with the specified name.
	
Examples:
  rc new my-app                    # Create a basic React app
  rc new my-app --ts              # Create a React app with TypeScript
  rc new my-app --tailwind        # Create a React app with Tailwind CSS
  rc new my-app --ts --tailwind   # Create a React app with TypeScript and Tailwind CSS`,
	Args: cobra.ExactArgs(1),
	RunE: runNewCommand,
}

// runNewCommand executes the new command logic
func runNewCommand(cmd *cobra.Command, args []string) error {
	appName := strings.TrimSpace(args[0])

	// Validate app name
	if err := validateAppName(appName); err != nil {
		return fmt.Errorf("invalid app name: %w", err)
	}

	// Create template data
	data := types.TemplateData{
		AppName:       appName,
		UseTypeScript: useTypeScript,
		UseTailwind:   useTailwind,
	}

	// Execute the new command
	if err := helpers.CreateReactApp(appName, data); err != nil {
		return fmt.Errorf("failed to create React app: %w", err)
	}

	return nil
}

// validateAppName validates the application name
func validateAppName(name string) error {
	if name == "" {
		return fmt.Errorf("app name cannot be empty")
	}

	// Check for invalid characters
	if strings.ContainsAny(name, "!@#$%^&*()+=[]{}|\\:;\"'<>?,./") {
		return fmt.Errorf("app name contains invalid characters")
	}

	// Check if name starts with a number
	if len(name) > 0 && name[0] >= '0' && name[0] <= '9' {
		return fmt.Errorf("app name cannot start with a number")
	}

	return nil
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Add flags with descriptions
	newCmd.Flags().BoolVarP(&useTypeScript, "ts", "t", false, "Use TypeScript instead of JavaScript")
	newCmd.Flags().BoolVarP(&useTailwind, "tailwind", "w", false, "Include Tailwind CSS in the project")
}
