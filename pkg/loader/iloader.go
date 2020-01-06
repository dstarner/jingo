package loader

import "html/template"

// ILoader represents a template loader
type ILoader interface {
	// GetTemplate returns a template given the path / name
	LoadTemplate(templateName string) (*template.Template, error)
}
