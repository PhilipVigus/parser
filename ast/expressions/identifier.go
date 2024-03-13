package expressions

import "lang/lexer/token"

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenValue() string {
	return i.Token.Value
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string {
	return i.Value
}
