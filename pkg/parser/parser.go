package parser

import (
	"io"

	"github.com/alecthomas/participle"
	"github.com/dstarner/jingo/pkg/parser/ast"
	"github.com/dstarner/jingo/pkg/parser/lexer"
	"github.com/pkg/errors"
)

// LanguageParser provides the functionality to break down a language into an AST
type LanguageParser struct {
	config       *parserConfig
	languageSpec *lexer.LanguageSpec
}

// NewLanguageParser creates a new parser given a list of parser and lexer options
func NewLanguageParser(parserOptions []Option, lexOptions []lexer.Option) *LanguageParser {
	config := newParserConfig(parserOptions...)
	return &LanguageParser{
		config:       config,
		languageSpec: lexer.NewLanguageSpec(lexOptions...),
	}
}

// ParseJinja runs Parse() and then converts the result into a JinjaAST
func (parser LanguageParser) ParseJinja(input io.Reader) (*ast.JinjaAST, error) {
	jinjaAST, err := parser.Parse(input)
	return jinjaAST.(*ast.JinjaAST), err
}

// Parse the input bytes into the interface designated by the parser configuration
func (parser LanguageParser) Parse(input io.Reader) (interface{}, error) {
	lexer, err := parser.languageSpec.Lexer()
	if err != nil {
		return nil, errors.Wrap(err, "error while loading parser")
	}
	internalParser, err := participle.Build(
		parser.config.astType,
		lexer,
		participle.Elide(parser.config.ignore...),
		participle.Unquote(parser.config.unquote...),
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while building parser")
	}
	err = internalParser.Parse(input, parser.config.astData)
	if err != nil {
		return nil, errors.Wrap(err, "error while running parser")
	}
	return parser.config.astData, nil
}
