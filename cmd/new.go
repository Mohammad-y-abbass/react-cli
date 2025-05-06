package cmd

import (
	"github.com/Mohammad-y-abbass/react-cli/helpers"
	"github.com/Mohammad-y-abbass/react-cli/types"
	"github.com/spf13/cobra"
)

var (
	useTs       bool
	useTailwind bool
)

var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new React app",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		data := types.TemplateData{
			UseTailwind: useTailwind,
		}

		appName := args[0]

		helpers.RunNewCmd(appName, useTs, useTailwind, data)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().BoolVarP(&useTs, "ts", "t", false, "Use TypeScript")
	newCmd.Flags().BoolVarP(&useTailwind, "tailwind", "w", false, "Use Tailwind CSS")
}
