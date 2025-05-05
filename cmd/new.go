package cmd

import (
	"os"
	"os/exec"

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

		appName := args[0]

		err := execViteCmd(appName)

		if err != nil {
			println("Error creating app:", err.Error())
			os.Exit(1)
		}

	},
}

func execViteCmd(appName string) error {

	var viteCmd *exec.Cmd

	if useTs {
		viteCmd = exec.Command("npm", "create", "vite@latest", appName, "--", "--template", "react-ts")
	} else {
		viteCmd = exec.Command("npm", "create", "vite@latest", appName, "--", "--template", "react")
	}
	viteCmd.Stdout = os.Stdout
	viteCmd.Stderr = os.Stderr

	return viteCmd.Run()
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().BoolVarP(&useTs, "ts", "t", false, "Use TypeScript")
}
