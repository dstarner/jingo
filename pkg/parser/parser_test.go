package parser_test

import (
	"strings"
	"testing"

	"github.com/dstarner/jingo/pkg/parser"
	"github.com/dstarner/jingo/pkg/parser/ast"
	"github.com/stretchr/testify/require"
)

func TestVariableParseAST(t *testing.T) {
	defaultConfiguration := map[string]string{
		"blockStartString":    "{%",
		"blockEndString":      "%}",
		"variableStartString": "{{",
		"variableEndString":   "}}",
	}

	given := "{#woohoo buddy #}{{ hello_world }}"
	expected := ast.JinjaAST{Variables: []ast.Variable{ast.Variable{Raw: &ast.Value{String: "hello_world"}}}}

	parser := parser.NewJinjaParser(defaultConfiguration)
	ast, err := parser.ParseAST(strings.NewReader(given))
	require.NoError(t, err)
	require.Equal(t, &expected, ast)
}
