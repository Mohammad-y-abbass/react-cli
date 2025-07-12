package types

// TemplateData holds the configuration data for project templates
type TemplateData struct {
	// UseTailwind indicates whether to include Tailwind CSS in the project
	UseTailwind bool
	// AppName is the name of the React application
	AppName string
	// UseTypeScript indicates whether to use TypeScript instead of JavaScript
	UseTypeScript bool
}
