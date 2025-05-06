package helpers

import (
	"os"
	"path/filepath"

	"github.com/Mohammad-y-abbass/react-cli/types"
)

func SetInitialCode(appName string, useTs bool, data types.TemplateData) {
	var appFilePath string
	appCssPath := filepath.Join(appName, "src", "App.css")

	if useTs {
		appFilePath = filepath.Join(appName, "src", "App.tsx")
	} else {
		appFilePath = filepath.Join(appName, "src", "App.jsx")
	}

	os.Remove(appCssPath)

	appCodeTmpl := filepath.Join("templates", "app.tpl")

	ApplyTemplate(appCodeTmpl, appFilePath, data)
}
