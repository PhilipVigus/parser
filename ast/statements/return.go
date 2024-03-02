package statements

import (
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
