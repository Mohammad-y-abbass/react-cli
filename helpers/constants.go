package helpers

// File permissions
const (
	DefaultFilePermission = 0644
)

// Template paths
const (
	AppTemplatePath            = "templates/app.tpl"
	ViteConfigTemplatePath     = "templates/vite-config.tpl"
	TailwindConfigTemplatePath = "templates/tailwind-config.tpl"
	PostCSSConfigTemplatePath  = "templates/postcss-config.tpl"
)

// File paths
const (
	SrcDir             = "src"
	AppCSSFile         = "App.css"
	IndexCSSFile       = "index.css"
	TailwindConfigFile = "tailwind.config.js"
	PostCSSConfigFile  = "postcss.config.js"
)

// Vite configuration
const (
	ReactTemplate     = "react"
	ReactTSTemplate   = "react-ts"
	TailwindCSSImport = `@tailwind base;
@tailwind components;
@tailwind utilities;`
)

// NPM commands
const (
	NPMCreateVite = "npm"
	NPMInstall    = "npm"
)

// Tailwind dependencies
var TailwindDependencies = []string{
	"tailwindcss",
	"postcss",
	"autoprefixer",
}

// Tailwind dev dependencies
var TailwindDevDependencies = []string{
	"@types/node",
}
