package statements

import (
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
