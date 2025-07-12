package helpers

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Mohammad-y-abbass/react-cli/types"
)

// CreateReactApp creates a new React application with the specified configuration
func CreateReactApp(appName string, data types.TemplateData) error {
	// Step 1: Create the Vite project
	if err := createViteProject(appName, data.UseTypeScript); err != nil {
		return fmt.Errorf("failed to create Vite project: %w", err)
	}

	// Step 2: Configure Tailwind CSS if requested
	if data.UseTailwind {
		if err := configureTailwindCSS(appName, data); err != nil {
			return fmt.Errorf("failed to configure Tailwind CSS: %w", err)
		}
	}

	// Step 3: Set up initial code
	if err := setupInitialCode(appName, data); err != nil {
		return fmt.Errorf("failed to setup initial code: %w", err)
	}

	// Step 4: Install dependencies
	if err := installDependencies(appName, data.UseTailwind); err != nil {
		return fmt.Errorf("failed to install dependencies: %w", err)
	}

	return nil
}

// createViteProject creates a new Vite project with React
func createViteProject(appName string, useTypeScript bool) error {
	GlobalLogger.Info("Creating Vite project...")

	template := ReactTemplate
	if useTypeScript {
		template = ReactTSTemplate
	}

	cmd := execCommand(NPMCreateVite, "create", "vite@latest", appName, "--", "--template", template)
	cmd.Stdout = io.Discard
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("vite command failed: %w", err)
	}

	GlobalLogger.Success("✓ Vite project created successfully")
	return nil
}

// configureTailwindCSS sets up Tailwind CSS in the project
func configureTailwindCSS(appName string, data types.TemplateData) error {
	GlobalLogger.Info("Configuring Tailwind CSS...")

	// Determine file paths based on TypeScript usage
	var viteConfigPath string
	if data.UseTypeScript {
		viteConfigPath = filepath.Join(appName, "vite.config.ts")
	} else {
		viteConfigPath = filepath.Join(appName, "vite.config.js")
	}

	indexCssPath := filepath.Join(appName, "src", "index.css")

	// Write Tailwind CSS import
	if err := writeFile(indexCssPath, []byte(TailwindCSSImport), DefaultFilePermission); err != nil {
		return fmt.Errorf("failed to write index.css: %w", err)
	}

	// Apply Vite configuration template
	viteConfigTmpl := ViteConfigTemplatePath
	if err := applyTemplate(viteConfigTmpl, viteConfigPath, data); err != nil {
		return fmt.Errorf("failed to apply Vite config template: %w", err)
	}

	GlobalLogger.Success("✓ Tailwind CSS configured successfully")
	return nil
}

// setupInitialCode sets up the initial React application code
func setupInitialCode(appName string, data types.TemplateData) error {
	GlobalLogger.Info("Setting up initial code...")

	// Determine app file path based on TypeScript usage
	var appFilePath string
	if data.UseTypeScript {
		appFilePath = filepath.Join(appName, "src", "App.tsx")
	} else {
		appFilePath = filepath.Join(appName, "src", "App.jsx")
	}

	// Remove default App.css
	appCssPath := filepath.Join(appName, SrcDir, AppCSSFile)
	if err := removeFile(appCssPath); err != nil {
		GlobalLogger.Warning("Failed to remove App.css: %v", err)
	}

	// Apply app code template
	appCodeTmpl := AppTemplatePath
	if err := applyTemplate(appCodeTmpl, appFilePath, data); err != nil {
		return fmt.Errorf("failed to apply app template: %w", err)
	}

	GlobalLogger.Success("✓ Initial code setup completed")
	return nil
}

// installDependencies installs project dependencies
func installDependencies(appName string, useTailwind bool) error {
	GlobalLogger.Info("Installing dependencies...")

	var cmd *exec.Cmd
	if useTailwind {
		args := append([]string{"install"}, TailwindDependencies...)
		cmd = execCommand(NPMInstall, args...)
	} else {
		cmd = execCommand(NPMInstall, "install")
	}

	cmd.Dir = appName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("npm install failed: %w", err)
	}

	GlobalLogger.Success("✓ Dependencies installed successfully")
	return nil
}
