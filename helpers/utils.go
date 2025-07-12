package helpers

import (
	"os"
	"os/exec"
	"text/template"

	"github.com/Mohammad-y-abbass/react-cli/types"
	"github.com/pterm/pterm"
)

// Logger provides structured logging functionality
type Logger struct{}

// Info logs an informational message
func (l *Logger) Info(message string) {
	pterm.Println(pterm.FgYellow.Sprint(message))
}

// Success logs a success message
func (l *Logger) Success(message string) {
	pterm.FgGreen.Println(message)
}

// Warning logs a warning message
func (l *Logger) Warning(format string, args ...interface{}) {
	pterm.Warning.Printf(format, args...)
}

// Error logs an error message
func (l *Logger) Error(message string) {
	pterm.Error.Println(message)
}

// Global logger instance
var GlobalLogger = &Logger{}

// execCommand creates a new exec.Cmd with the given command and arguments
func execCommand(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}

// writeFile writes data to a file with the specified permissions
func writeFile(path string, data []byte, perm os.FileMode) error {
	return os.WriteFile(path, data, perm)
}

// removeFile removes a file if it exists
func removeFile(path string) error {
	return os.Remove(path)
}

// applyTemplate applies a template to a file
func applyTemplate(templatePath, outputPath string, data types.TemplateData) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
