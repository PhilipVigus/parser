package token

import "fmt"

type Type int
type Token[T any] struct {
	Type    Type
	Literal T
}

func New[T any](t Type, l T) Token[any] {
	return Token[any]{
		Type:    t,
		Literal: l,
	}
}

func (t Token[T]) String() string {
	return fmt.Sprintf("[ type: %s, literal: %v ]", GetTokenType(t), t.Literal)
}

const (
	ILLEGAL = iota
	EOF
	IDENT
	ASSIGN
	PLUS
	MINUS
	DIVIDE
	MULTIPLY
	COMMA
	SEMICOLON
	COLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	FUNCTION
	LET
)

var tokenTypes = map[Type]string{
	ILLEGAL:   "ILLEGAL",
	EOF:       "EOF",
	IDENT:     "IDENT",
	ASSIGN:    "ASSIGN",
	PLUS:      "PLUS",
	MINUS:     "MINUS",
	DIVIDE:    "DIVIDE",
	MULTIPLY:  "MULTIPLY",
	COMMA:     "COMMA",
	SEMICOLON: "SEMICOLON",
	COLON:     "COLON",
	LPAREN:    "LPAREN",
	RPAREN:    "RPAREN",
	LBRACE:    "LBRACE",
	RBRACE:    "RBRACE",
	FUNCTION:  "FUNCTION",
	LET:       "LET",
}

func GetTokenType[T any](t Token[T]) string {
	name, exists := tokenTypes[t.Type]

	if exists {
		return name
	}

	return "UNKNOWN"
}
