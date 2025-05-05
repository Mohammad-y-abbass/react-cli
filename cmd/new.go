package cmd

import (
	"os"
	"os/exec"
	"path/filepath"

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

		err := execViteCmd(appName)

		if useTailwind {
			addTailwindConfig(appName, data)

		}

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

func addTailwindConfig(appName string, data types.TemplateData) error {

	var viteConfigPath string

	if useTs {
		viteConfigPath = filepath.Join(appName, "vite.config.ts")
	} else {
		viteConfigPath = filepath.Join(appName, "vite.config.js")
	}

	viteConfigTmpl := filepath.Join("templates", "vite-config.tpl")

	helpers.ApplyTemplate(viteConfigTmpl, viteConfigPath, data)

	return nil

}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().BoolVarP(&useTs, "ts", "t", false, "Use TypeScript")
	newCmd.Flags().BoolVarP(&useTailwind, "tailwind", "w", false, "Use Tailwind CSS")
}
