package statements

import (
	"lang/ast/expressions"
	"lang/lexer/token"
)

type Define struct {
	Token token.Token
	Name  *expressions.Identifier
	Value expressions.Expression
}

func (a *Define) TokenValue() string {
	return a.Token.Value
}

func (a *Define) statementNode() {}
