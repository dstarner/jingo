package ast

// Represents an Operator term such as +, -, /, //...
type Operator int

const (
	OpMult Operator = iota
	OpFloorDiv
	OpDiv
	OpAdd
	OpSub
)

var operatorMap = map[string]Operator{
	"+":  OpAdd,
	"-":  OpSub,
	"*":  OpMult,
	"/":  OpDiv,
	"//": OpFloorDiv,
}

// Capture allows us to make an operator value to a constant value
func (o *Operator) Capture(s []string) error {
	*o = operatorMap[s[0]]
	return nil
}

// JinjaAST represents a Jinja Abstract Syntax Tree that the given input gets parsed into
type JinjaAST struct {
}

type Value struct {
	Number        *float64    `  @(Float|Int)`
	Variable      *string     `| @Ident`
	Subexpression *Expression `| "(" @@ ")"`
}

type Factor struct {
	Base     *Value `@@`
	Exponent *Value `[ "^" @@ ]`
}

type OpFactor struct {
	Operator Operator `@("*" | "/")`
	Factor   *Factor  `@@`
}

type Term struct {
	Left  *Factor     `@@`
	Right []*OpFactor `{ @@ }`
}

type OpTerm struct {
	Operator Operator `@("+" | "-")`
	Term     *Term    `@@`
}

type Expression struct {
	Left  *Term     `@@`
	Right []*OpTerm `{ @@ }`
}
