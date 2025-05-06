package helpers

import (
	"io"
	"os"
	"os/exec"

	"github.com/pterm/pterm"
)

func ExecViteCmd(appName string, useTs bool) {
	pterm.Println(pterm.FgYellow.Sprint("Setting up project..."))

	var viteCmd *exec.Cmd

	if useTs {
		viteCmd = exec.Command("npm", "create", "vite@latest", appName, "--", "--template", "react-ts")
	} else {
		viteCmd = exec.Command("npm", "create", "vite@latest", appName, "--", "--template", "react")
	}
	viteCmd.Stdout = io.Discard
	viteCmd.Stderr = os.Stderr

	viteCmd.Run()
	pterm.FgGreen.Println("✔️  Done")
}
