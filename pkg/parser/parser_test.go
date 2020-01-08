package parser_test

import (
	"strings"
	"testing"

	"github.com/alecthomas/repr"

	"github.com/dstarner/jingo/pkg/parser"
	"github.com/dstarner/jingo/pkg/parser/lexer"
	"github.com/stretchr/testify/require"
)

// Shamelessly stolen from the following github repository because I needed to test the
// parsing configuration and management worked before I knew how to write a Jinja grammar block
// https://github.com/electricjesus/nginx-ebnf/
func TestNGINXConfigParse(t *testing.T) {
	languageTemplate := `
	Ident = (alpha | "_") { "_" | alpha | digit } 					.
	Comment = ("#") { "\u0000"…"\uffff"-"\n" } ("\n")				.
	Number = ("." | digit) {"." | digit} 							.
	String = ("\"") { "\u0000"…"\uffff"-"\n"-"\"" } ("\"")			.
	KeyVal = { Ident } ("=") { alpha | digit }						.
	Whitespace = " " | "\t" | "\n" | "\r" 							.
	Punct = "!"…"/" | ":"…"@" | "["…` + "\"`\"" + ` | "{"…"~" 		.
	alpha = "a"…"z" | "A"…"Z" 										.
	digit = "0"…"9" 												.
	`

	// Block - Block recursive structure
	type block struct {
		// TODO: Pos       lexer.Position
		Children  []*block `parser:"(\"{\" @@* \"}\")?" json:"children,omitempty"`
		Directive string   `parser:"@Ident?" json:"directive,omitempty"`
		Args      []string `parser:"(@(Ident | Number | String | KeyVal)* \";\")?" json:"args,omitempty"`
		Comment   string   `parser:"(@Comment)?" json:"comment,omitempty"`
	}

	expected := &block{
		Children: []*block{
			&block{Comment: " a comment"},
			&block{Comment: " comment2"},
			&block{Directive: "http"},
			&block{Children: []*block{
				&block{Comment: " another comment"},
				&block{Directive: "server"},
				&block{Children: []*block{
					&block{Directive: "listen", Args: []string{"127.0.0.1"}},
				},
				},
			},
			},
		},
	}

	given := `
	{
		# a comment
		# comment2
		http {
			# another comment
			server {
				listen 127.0.0.1;
			}
		}
	}`
	actual := &block{}

	languageParser := parser.NewLanguageParser(
		[]parser.Option{
			parser.ASTType(&block{}),
			parser.ASTData(actual),
			parser.Ignore("Whitespace"),
			parser.Unquote("String", "Comment"),
		},
		[]lexer.Option{
			lexer.LanguageTemplate(languageTemplate),
		},
	)
	_, err := languageParser.Parse(strings.NewReader(given))
	require.NoError(t, err)
	repr.Println(actual, repr.IgnoreGoStringer())

	require.Equal(t, expected, actual)
}
