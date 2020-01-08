package parser

import "github.com/dstarner/jingo/pkg/parser/ast"

// Option represents a mutator function that changes a ParserConfig
type Option func(*parserConfig)

// -------
// Options
// -------

// ASTType is an empty struct pointer of the AST struct used to build the parser
func ASTType(astType interface{}) Option {
	return func(config *parserConfig) {
		config.astType = astType
	}
}

// ASTData is an empty struct pointer of the AST struct that data will be saved to
func ASTData(astData interface{}) Option {
	return func(config *parserConfig) {
		config.astData = astData
	}
}

// Ignore certain identifiers from the regex
func Ignore(ignorable ...string) Option {
	return func(config *parserConfig) {
		config.ignore = ignorable
	}
}

// Unquote certain identifiers from the regex
func Unquote(unquote ...string) Option {
	return func(config *parserConfig) {
		config.unquote = unquote
	}
}

type parserConfig struct {
	ignore  []string
	unquote []string
	astType interface{}
	astData interface{}
}

func newParserConfig(options ...Option) *parserConfig {
	config := &parserConfig{
		ignore:  []string{"Comment", "BlockComment"},
		unquote: []string{},
		astType: &ast.JinjaAST{},
		astData: &ast.JinjaAST{},
	}
	config.Configure(options...)
	return config
}

// Configure the SpecConfig after initial creation
func (config *parserConfig) Configure(options ...Option) {
	for _, option := range options {
		option(config)
	}
}
