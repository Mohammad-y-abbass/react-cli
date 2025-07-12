package helpers

import (
	"bytes"
	"os"
	"text/template"

	"github.com/Mohammad-y-abbass/react-cli/types"
)

func ApplyTemplate(fileTmpl string, filePath string, data types.TemplateData) {
	tmpl, err := template.ParseFiles(fileTmpl)
	if err != nil {
		GlobalLogger.Error("Error parsing template: " + err.Error())
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		GlobalLogger.Error("Error executing template: " + err.Error())
	}

	err = os.WriteFile(filePath, buffer.Bytes(), 0644)
	if err != nil {
		GlobalLogger.Error("Error writing file: " + err.Error())
	}

}
