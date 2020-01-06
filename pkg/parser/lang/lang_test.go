package lang_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/regex"
	"github.com/alecthomas/repr"
	"github.com/dstarner/jingo/pkg/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBasicLang(t *testing.T) {

	defaultConfiguration := map[string]string{
		"blockStartString":    "{%",
		"blockEndString":      "%}",
		"variableStartString": "{{",
		"variableEndString":   "}}",
	}
	language := parser.GenerateLanguageRegex(defaultConfiguration)
	d, err := regex.New(language)

	symbols := d.Symbols()
	require.NoError(t, err)

	testCases := []struct {
		source   string
		expected []lexer.Token
	}{
		{
			source: "{{}}",
			expected: []lexer.Token{
				lexer.Token{Type: symbols["VariableOpen"], Value: "{{", Pos: lexer.Position{Line: 1, Column: 1}},
				lexer.Token{Type: symbols["VariableClose"], Value: "}}", Pos: lexer.Position{Offset: 2, Line: 1, Column: 3}},
			},
		},
		{
			source: "{%%}",
			expected: []lexer.Token{
				lexer.Token{Type: symbols["BlockOpen"], Value: "{%", Pos: lexer.Position{Line: 1, Column: 1}},
				lexer.Token{Type: symbols["BlockClose"], Value: "%}", Pos: lexer.Position{Offset: 2, Line: 1, Column: 3}},
			},
		},
	}

	for _, test := range testCases {
		l, err := d.Lex(strings.NewReader(test.source))
		require.NoError(t, err)

		actual, err := lexer.ConsumeAll(l)
		require.NoError(t, err)

		actualStrip := actual[:len(actual)-1]

		if !assert.Equal(t, test.expected, actualStrip) {
			fmt.Println(language)
			repr.Println(symbols, repr.IgnoreGoStringer())
			fmt.Println()

			fmt.Println("Actual:")
			repr.Println(actualStrip, repr.IgnoreGoStringer())
			fmt.Println("Expected:")
			repr.Println(test.expected, repr.IgnoreGoStringer())
			repr.Println("------------")
			t.FailNow()
		}
	}
}
