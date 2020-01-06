package lang

import (
	"fmt"

	"github.com/dstarner/jingo/pkg/strutil"
)

// GenerateLanguageRegex generates the language specs for Jinja given a map of options for variable terms
func GenerateLanguageRegex(options interface{}) string {
	res, err := strutil.StringFromTemplate(jinjaLanguageRegex, options)
	if err != nil {
		panic(fmt.Sprintf("Incorrect options passed to GenerateLanguageRegex: %v", err))
	}
	return res
}

// JinjaLanguageRegex represents the grammar for Jinja
const jinjaLanguageRegex = `
BlockComment = blockCommentStartblockCommentEnd

Identity = [[:alpha:]]\w*
Number = ("." | digit) {"." | digit} .


VariableOpen  = {{ .variableStartString }}[ ]*
VariableClose = [ ]*{{ .variableEndString }}
BlockOpen     = {{ .blockStartString }}[ ]*
BlockClose    = [ ]*{{ .blockEndString }}

blockCommentStart = {#
blockCommentEnd   = #}

alpha = "a"…"z" | "A"…"Z" .
digit = "0"…"9" .


Whitespace    = [ \t\n\r] 
`

// TokenFloorDiv = //
// TokenDiv      = /
// TokenPow      = \*\*
// TokenMul      = \*

// TokenAdd      = \+
// TokenSub      = -
// TokenMod      = %

// TokenEqual    = ==
// TokenNotEqual = !=
// TokenGTEqual  = >=
// TokenLTEqual  = <=

// TokenTilde    = ~
// TokenAssign   = =
// TokenColon    = :
// TokenComma    = ,

// TokenLBracket = \[
// TokenRBracket = \]
// TokenLParen   = \(
// TokenRParen   = \)
// TokenLBrace   = {
// TokenRBrace   = }
// `
