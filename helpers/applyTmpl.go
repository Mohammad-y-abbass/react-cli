package helpers

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/Mohammad-y-abbass/react-cli/types"
)

func ApplyTemplate(fileTmpl string, filePath string, data types.TemplateData) {
	tmpl, err := template.ParseFiles(fileTmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}

	err = os.WriteFile(filePath, buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}

}
