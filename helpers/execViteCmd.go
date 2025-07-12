package helpers

import (
	"io"
	"os"
	"os/exec"
)

func ExecViteCmd(appName string, useTs bool) {
	GlobalLogger.Info("Setting up project...")

	var viteCmd *exec.Cmd

	if useTs {
		viteCmd = exec.Command("npm", "create", "vite@latest", appName, "--", "--template", "react-ts")
	} else {
		viteCmd = exec.Command("npm", "create", "vite@latest", appName, "--", "--template", "react")
	}
	viteCmd.Stdout = io.Discard
	viteCmd.Stderr = os.Stderr

	viteCmd.Run()
	GlobalLogger.Success("✔️  Done")
}
