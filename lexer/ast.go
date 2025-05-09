package rythin

type Program struct {
	Functions []FunctionDef
}

type FunctionDef struct {
	Name       string
	Params     []Param
	ReturnType string
	Body       []Statement
}

type Param struct {
	Name string
	Type string
}

type Statement interface{}

type ReturnStatement struct {
	Value Expression
}

type Expression interface{}

type NumberLiteral struct {
	Value string
}

type Identifier struct {
	Name string
}
