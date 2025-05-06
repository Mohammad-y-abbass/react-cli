package helpers

import (
	"os"
	"os/exec"
)

func ExecViteCmd(appName string, useTs bool) {

	var viteCmd *exec.Cmd

	if useTs {
		viteCmd = exec.Command("npm", "create", "vite@latest", appName, "--", "--template", "react-ts")
	} else {
		viteCmd = exec.Command("npm", "create", "vite@latest", appName, "--", "--template", "react")
	}
	viteCmd.Stdout = os.Stdout
	viteCmd.Stderr = os.Stderr

	viteCmd.Run()
}
