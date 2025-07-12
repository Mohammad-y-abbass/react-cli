package helpers

// ProjectConfig holds configuration for the React CLI project
type ProjectConfig struct {
	// Version of the CLI tool
	Version string
	// Default template to use
	DefaultTemplate string
	// Supported templates
	SupportedTemplates []string
}

// DefaultConfig returns the default configuration
func DefaultConfig() *ProjectConfig {
	return &ProjectConfig{
		Version:            "1.0.0",
		DefaultTemplate:    ReactTemplate,
		SupportedTemplates: []string{ReactTemplate, ReactTSTemplate},
	}
}

// IsValidTemplate checks if the given template is supported
func (c *ProjectConfig) IsValidTemplate(template string) bool {
	for _, supported := range c.SupportedTemplates {
		if supported == template {
			return true
		}
	}
	return false
}
