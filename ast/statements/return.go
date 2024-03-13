package statements

import (
	"bytes"
	"lang/ast/expressions"
	"lang/lexer/token"
)

type Return struct {
	Token       token.Token
	ReturnValue expressions.Expression
}

func (r *Return) TokenValue() string {
	return r.Token.Value
}

func (r *Return) statementNode() {}

func (r *Return) String() string {
	var out bytes.Buffer

	out.WriteString(r.TokenValue() + " ")

	if r.ReturnValue != nil {
		out.WriteString(r.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
