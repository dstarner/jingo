package parser

import (
	"io"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/regex"
	"github.com/dstarner/jingo/pkg/parser/ast"
	"github.com/dstarner/jingo/pkg/parser/lang"
)

type JinjaParser struct {
	jinjaLang       string
	jinjaDefinition lexer.Definition
}

func NewJinjaParser(config interface{}) *JinjaParser {
	lang := lang.GenerateLanguageRegex(config)
	return &JinjaParser{
		jinjaLang:       lang,
		jinjaDefinition: lexer.Must(regex.New(lang)),
	}
}

func (parser JinjaParser) ParseAST(input io.Reader) (*ast.JinjaAST, error) {
	parsedAST := &ast.JinjaAST{}
	participleParser := participle.MustBuild(&ast.JinjaAST{},
		participle.Lexer(parser.jinjaDefinition),
		participle.Elide("Comment"),
	)
	err := participleParser.Parse(input, parsedAST)
	return parsedAST, err
}
