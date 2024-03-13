package statements

import (
	"bytes"
	"lang/ast/expressions"
	"lang/lexer/token"
)

type Assign struct {
	Token token.Token
	Name  *expressions.Identifier
	Value expressions.Expression
}

func (a *Assign) TokenValue() string {
	return a.Token.Value
}

func (a *Assign) statementNode() {}

func (a *Assign) String() string {
	var out bytes.Buffer

	out.WriteString(a.TokenValue() + " ")
	out.WriteString(a.Name.String())
	out.WriteString(" = ")

	if a.Value != nil {
		out.WriteString(a.Value.String())
	}
	out.WriteString(";")

	return out.String()
}
