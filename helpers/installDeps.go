package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallDeps(appName string, useTailwind bool) {
	var npmCmd *exec.Cmd
	if useTailwind {
		npmCmd = exec.Command("npm", "install", "tailwindcss", "@tailwindcss/vite")
	} else {
		npmCmd = exec.Command("npm", "install")
	}
	npmCmd.Dir = appName
	npmCmd.Stdout = os.Stdout
	npmCmd.Stderr = os.Stderr

	if err := npmCmd.Run(); err != nil {
		fmt.Println("Error running npm install:", err)
		os.Exit(1)
	}
}
