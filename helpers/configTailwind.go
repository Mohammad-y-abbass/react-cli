package helpers

import (
	"os"
	"path/filepath"

	"github.com/Mohammad-y-abbass/react-cli/types"
)

func ConfigTailwind(appName string, data types.TemplateData, useTs bool) {
	GlobalLogger.Info("Setting up tailwind css...")

	var viteConfigPath string

	if useTs {
		viteConfigPath = filepath.Join(appName, "vite.config.ts")
	} else {
		viteConfigPath = filepath.Join(appName, "vite.config.js")
	}

	indexCssPath := filepath.Join(appName, "src", "index.css")

	os.WriteFile(indexCssPath, []byte(`@import "tailwindcss";`), 0644)

	viteConfigTmpl := filepath.Join("templates", "vite-config.tpl")

	ApplyTemplate(viteConfigTmpl, viteConfigPath, data)

	GlobalLogger.Success("✔️  Done")

}
