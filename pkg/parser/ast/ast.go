package ast

type JinjaAST struct {
	Variables []Variable `@@`
}

type Variable struct {
	Raw *Value `VariableOpen @@ VariableClose`
}

type Value struct {
	String string `@Identity`
}
