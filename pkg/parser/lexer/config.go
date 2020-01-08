package lexer

import (
	"github.com/dstarner/jingo/pkg/parser/lexer/templates"
	"github.com/dstarner/jingo/pkg/strutil"
	"github.com/pkg/errors"
)

// Option allows configuring the SpecConfig
type Option func(*SpecConfig)

// -------
// Options
// -------

// BlockStartIdentifier sets the substring that identifies when a block statement starts
func BlockStartIdentifier(identifier string) Option {
	return func(config *SpecConfig) {
		config.blockStartIdentifier = identifier
	}
}

// BlockEndIdentifier sets the substring that identifies when a block statement ends
func BlockEndIdentifier(identifier string) Option {
	return func(config *SpecConfig) {
		config.blockEndIdentifier = identifier
	}
}

// VariableStartIdentifier sets the substring that identifies when a variable starts
func VariableStartIdentifier(identifier string) Option {
	return func(config *SpecConfig) {
		config.variableStartIdentifier = identifier
	}
}

// VariableEndIdentifier sets the substring that identifies when a variable ends
func VariableEndIdentifier(identifier string) Option {
	return func(config *SpecConfig) {
		config.variableEndIdentifier = identifier
	}
}

// LanguageTemplate sets the core language regex to use, as defined by
// https://github.com/alecthomas/participle/blob/master/lexer/regex/regex.go#L24
// It also can contain references to any variables in the SpecConfig
func LanguageTemplate(template string) Option {
	return func(config *SpecConfig) {
		config.languageTemplate = template
	}
}

// --------------------
// Configuration Struct
// --------------------

// SpecConfig represents language configuration that can be
// used to build the language parser
type SpecConfig struct {
	blockStartIdentifier string
	blockEndIdentifier   string

	variableStartIdentifier string
	variableEndIdentifier   string

	blockCommentStartIdentifier string
	blockCommentEndIdentifier   string

	languageTemplate string
}

// NewSpecConfig creates a new default language specification configuration
func NewSpecConfig(options ...Option) *SpecConfig {
	config := &SpecConfig{
		blockStartIdentifier:        "{%",
		blockEndIdentifier:          "%}",
		variableStartIdentifier:     "{{",
		variableEndIdentifier:       "}}",
		blockCommentStartIdentifier: "{#",
		blockCommentEndIdentifier:   "#}",
		languageTemplate:            templates.JinjaLanguageTemplate,
	}
	config.Configure(options...)
	return config
}

// Configure the SpecConfig after initial creation
func (config *SpecConfig) Configure(options ...Option) {
	for _, option := range options {
		option(config)
	}
}

// GenerateRegexString templates the configuration into the language spec
func (config SpecConfig) GenerateRegexString() (string, error) {
	regexStr, err := strutil.StringFromTemplate(config.languageTemplate, config.Map())
	if err != nil {
		return "", errors.Wrap(err, "could not generate language grammar")
	}
	return regexStr, nil
}

// Map can probably be rewritten with reflection to map it easier
func (config SpecConfig) Map() map[string]interface{} {
	return map[string]interface{}{
		"blockStartIdentifier": config.blockStartIdentifier,
		"blockEndIdentifier":   config.blockEndIdentifier,

		"variableStartIdentifier": config.variableStartIdentifier,
		"variableEndIdentifier":   config.variableEndIdentifier,

		"blockCommentStartIdentifier": config.blockCommentStartIdentifier,
		"blockCommentEndIdentifier":   config.blockCommentEndIdentifier,
	}
}
