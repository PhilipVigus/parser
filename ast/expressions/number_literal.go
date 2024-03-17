package expressions

import "lang/lexer/token"

type Number interface {
	int64 | float64
}

type NumberLiteral[T Number] struct {
	Token token.Token
	Value T
}

func (nl *NumberLiteral[T]) TokenValue() string {
	return nl.Token.Value
}

func (nl *NumberLiteral[T]) expressionNode() {}

func (nl *NumberLiteral[T]) String() string {
	return nl.Token.Value
}
