package helpers

import (
	"os"
	"os/exec"
)

func InstallDeps(appName string, useTailwind bool) {

	GlobalLogger.Info("Installing dependencies...")

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
		GlobalLogger.Error("Error running npm install: " + err.Error())
		os.Exit(1)
	}

	GlobalLogger.Success("✔️  Done")
}
