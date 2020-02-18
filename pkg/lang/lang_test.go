package lang_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/regex"
	"github.com/alecthomas/repr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/dstarner/jingo/pkg/lang"
)

func TestBasicLang(t *testing.T) {

	defaultConfiguration := map[string]string{
		"blockStartString":    "{%",
		"blockEndString":      "%}",
		"variableStartString": "{{",
		"variableEndString":   "}}",
	}
	language := lang.GenerateLanguageRegex(defaultConfiguration)
	d, err := regex.New(language)

	symbols := d.Symbols()
	require.NoError(t, err)

	testCases := []struct {
		source   string
		expected []lexer.Token
	}{
		// TEST WITH SOME VARIABLES
		// ------------------------
		{
			source: "{{}}",
			expected: []lexer.Token{
				lexer.Token{Type: symbols["VariableOpen"], Value: "{{"},
				lexer.Token{Type: symbols["VariableClose"], Value: "}}"},
			},
		},
		// TEST WITH SOME BLOCKS
		// ---------------------
		{
			source: "{%%}",
			expected: []lexer.Token{
				lexer.Token{Type: symbols["BlockOpen"], Value: "{%"},
				lexer.Token{Type: symbols["BlockClose"], Value: "%}"},
			},
		},
		// TEST WITH SOME NON-WRAPPED TEXT AND THEN A BLOCK
		// ------------------------------------------------
		{
			source: "some text {% if %}",
			expected: []lexer.Token{
				lexer.Token{ Type: symbols["Identity"], Value: "some"},
				lexer.Token{ Type: symbols["Whitespace"], Value: " "},
				lexer.Token{ Type: symbols["Identity"], Value: "text"},
				lexer.Token{ Type: symbols["Whitespace"], Value: " "},
				lexer.Token{ Type: symbols["BlockOpen"], Value: "{% "},
				lexer.Token{ Type: symbols["Identity"], Value: "if"},
				lexer.Token{ Type: symbols["BlockClose"], Value: " %}"},
			  },
		},
		// TEST WITH A METHOD PIPE
		// -----------------------
		{
			source: "{{ var_name|lower }}",
			expected: []lexer.Token{
				lexer.Token{ Type: symbols["VariableOpen"], Value: "{{ "},
				lexer.Token{ Type: symbols["Identity"], Value: "var_name"},
				lexer.Token{ Type: symbols["Pipe"], Value: "|"},
				lexer.Token{ Type: symbols["Identity"], Value: "lower"},
				lexer.Token{ Type: symbols["VariableClose"], Value: " }}"},
			  },
		},
		{
			source: "{{ my_age|add(5)|sub(-5.0)|mult(.03) }}",
			expected: []lexer.Token{
				lexer.Token{ Type: symbols["VariableOpen"], Value: "{{ " },
				lexer.Token{ Type: symbols["Identity"], Value: "my_age" },
				lexer.Token{ Type: symbols["Pipe"], Value: "|" },
				lexer.Token{ Type: symbols["Identity"], Value: "add" },
				lexer.Token{ Type: symbols["OpenParen"], Value: "(" },
				lexer.Token{ Type: symbols["Number"], Value: "5" },
				lexer.Token{ Type: symbols["CloseParen"], Value: ")" },
				lexer.Token{ Type: symbols["Pipe"], Value: "|" },
				lexer.Token{ Type: symbols["Identity"], Value: "sub" },
				lexer.Token{ Type: symbols["OpenParen"], Value: "(" },
				lexer.Token{ Type: symbols["Number"], Value: "-5.0" },
				lexer.Token{ Type: symbols["CloseParen"], Value: ")" },
				lexer.Token{ Type: symbols["Pipe"], Value: "|" },
				lexer.Token{ Type: symbols["Identity"], Value: "mult" },
				lexer.Token{ Type: symbols["OpenParen"], Value: "(" },
				lexer.Token{ Type: symbols["Number"], Value: ".03" },
				lexer.Token{ Type: symbols["CloseParen"], Value: ")" },
				lexer.Token{ Type: -6, Value: " }}" },
			  },
		},
	}

	for _, test := range testCases {
		l, err := d.Lex(strings.NewReader(test.source))
		require.NoError(t, err)

		actual, err := lexer.ConsumeAll(l)
		require.NoError(t, err)

		testExpected := []lexer.Token{}
		for _, token := range test.expected {
			token.Pos = lexer.Position{}
			testExpected = append(testExpected, token)
		}

		actualStrip := []lexer.Token{}
		for _, token := range actual[:len(actual)-1] {
			token.Pos = lexer.Position{}
			actualStrip = append(actualStrip, token)
		}

		if !assert.Equal(t, testExpected, actualStrip) {
			fmt.Println("Actual:")
			repr.Println(actualStrip, repr.IgnoreGoStringer())
			fmt.Println("Expected:")
			repr.Println(testExpected, repr.IgnoreGoStringer())
			repr.Println("------------")
			t.FailNow()
		}
	}
}
