package lexer_test

import (
	"testing"

	"github.com/dstarner/jingo/pkg/parser/lexer"
	"github.com/stretchr/testify/require"
)

func TestGenerateRegexString(t *testing.T) {
	variableOpen := "{{"
	template := `
	VariableOpen = {{ .variableStartIdentifier }}
	`
	expected := `
	VariableOpen = {{
	`

	config := lexer.NewSpecConfig(
		lexer.VariableEndIdentifier("}}"),
		lexer.BlockEndIdentifier("%}"),
		lexer.BlockStartIdentifier("{%"),
		lexer.VariableStartIdentifier(variableOpen),
		lexer.LanguageTemplate(template),
	)
	actual, err := config.GenerateRegexString()
	require.NoError(t, err)

	require.Equal(t, expected, actual)
}

func TestInvalidGoTemplate(t *testing.T) {
	template := `
	VariableOpen = {{  }}
	`

	config := lexer.NewSpecConfig(lexer.LanguageTemplate(template))
	_, err := config.GenerateRegexString()
	require.Error(t, err)
}
