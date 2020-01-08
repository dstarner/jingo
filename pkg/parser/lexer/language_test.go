package lexer_test

import (
	"strings"
	"testing"

	"github.com/alecthomas/repr"

	participleLexer "github.com/alecthomas/participle/lexer"
	"github.com/dstarner/jingo/pkg/parser/lexer"
	"github.com/stretchr/testify/require"
)

func TestInvalidGenerateDefinition(t *testing.T) {
	template := `
	VariableOpen "{{ .variableStartIdentifier }}" .
	`

	spec := lexer.NewLanguageSpec(lexer.LanguageTemplate(template))
	_, err := spec.GenerateDefinition()
	require.Error(t, err)
}

func TestTokenize(t *testing.T) {
	template := `
VariableOpen = "{{ .variableStartIdentifier }}" .
VariableClose = "{{ .variableEndIdentifier }}" .
	`
	expected := []participleLexer.Token{
		participleLexer.Token{Type: -2, Value: "{{", Pos: participleLexer.Position{Line: 1, Column: 1}},
		participleLexer.Token{Type: -3, Value: "}}", Pos: participleLexer.Position{Offset: 2, Line: 1, Column: 3}},
		participleLexer.Token{Type: -1, Pos: participleLexer.Position{Offset: 4, Line: 1, Column: 5}},
	}

	spec := lexer.NewLanguageSpec(lexer.LanguageTemplate(template))
	tokens, err := spec.Tokenize(strings.NewReader("{{}}"))
	require.NoError(t, err)

	repr.Println(tokens, repr.IgnoreGoStringer())
	require.Equal(t, expected, tokens)
}

func TestSymbolTypes(t *testing.T) {
	template := `
VariableOpen = "{{ .variableStartIdentifier }}" .
VariableClose = "{{ .variableEndIdentifier }}" .
	`
	expected := map[string]int32{
		"EOF":           -1,
		"VariableOpen":  -2,
		"VariableClose": -3,
	}

	spec := lexer.NewLanguageSpec(lexer.LanguageTemplate(template))
	actual, err := spec.SymbolTypes()
	require.NoError(t, err)

	repr.Println(actual, repr.IgnoreGoStringer())
	require.Equal(t, expected, actual)
}
