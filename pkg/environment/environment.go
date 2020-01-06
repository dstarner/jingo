package environment

import (
	"errors"
	"html/template"

	"github.com/dstarner/jingo/pkg/loader"
)

// Environment is the core component of Jinja. It contains important shared
// variables like configuration, filters, tests, globals and others.
type Environment struct {
	configuration  *Configuration
	templateLoader loader.ILoader
}

// NewEnvironment returns a new environment with a configuration built by the given options
func NewEnvironment(templateLoader loader.ILoader, options ...Option) *Environment {
	configuration := NewConfiguration()
	for _, opt := range options {
		opt(configuration)
	}

	environment := &Environment{
		configuration:  configuration,
		templateLoader: templateLoader,
	}
	return environment
}

func (env Environment) loadTemplate(templateName string) (*template.Template, error) {
	if env.templateLoader == nil {
		return nil, errors.New("no loader for this environment specified")
	}
	// TODO: add template cache here later on

	return env.templateLoader.LoadTemplate(templateName)
}

func (env Environment) FromString(source string) string {
	return ""
}
