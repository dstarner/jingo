package lexer

import (
	"io"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
	"github.com/pkg/errors"
)

// LanguageSpec creates a language definition and basic lexer for the parser to use
type LanguageSpec struct {
	config *SpecConfig
}

// NewLanguageSpec creates a new language spec from the options potentially given
func NewLanguageSpec(options ...Option) *LanguageSpec {
	return &LanguageSpec{
		config: NewSpecConfig(options...),
	}
}

// GenerateDefinition creates the lexer definition to use when parsing
func (spec LanguageSpec) GenerateDefinition() (lexer.Definition, error) {
	errWrap := func(err error) error {
		return errors.Wrap(err, "could not generate language definition")
	}

	regexStr, err := spec.config.GenerateRegexString()
	if err != nil {
		return nil, errWrap(err)
	}
	definition, err := ebnf.New(regexStr)
	if err != nil {
		return nil, errWrap(err)
	}
	return definition, nil
}

// Lexer returns a participle.Option that will set the Lexer for a Parser
func (spec LanguageSpec) Lexer() (participle.Option, error) {
	definition, err := spec.GenerateDefinition()
	if err != nil {
		return nil, errors.Wrap(err, "could not create lexer")
	}
	return participle.Lexer(definition), nil
}

// Tokenize returns a list of tokens after lex-ing the given input Reader
func (spec LanguageSpec) Tokenize(input io.Reader) ([]lexer.Token, error) {
	errWrap := func(err error) error {
		return errors.Wrap(err, "could not tokenize")
	}

	definition, err := spec.GenerateDefinition()
	if err != nil {
		return nil, errWrap(err)
	}
	l, err := definition.Lex(input)
	if err != nil {
		return nil, errWrap(err)
	}
	tokens, err := lexer.ConsumeAll(l)
	if err != nil {
		return nil, errWrap(err)
	}
	return tokens, nil
}

// SymbolTypes returns a map of symbol type name mapped to its internal ID
func (spec LanguageSpec) SymbolTypes() (map[string]rune, error) {
	definition, err := spec.GenerateDefinition()
	if err != nil {
		return nil, errors.Wrap(err, "could not generate symbol types")
	}
	return definition.Symbols(), nil
}
