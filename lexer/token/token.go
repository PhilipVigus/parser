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
	Illegal = iota
	Eof
	Ident
	Assign
	Plus
	Minus
	Divide
	Multiply
	Modulus
	Comma
	FullStop
	Semicolon
	Colon
	LParen
	RParen
	LBrace
	RBrace
	LBracket
	RBracket
	Percent
	DoubleQuote
	SingleQuote
	GreaterThan
	LessThan
	GreaterThanOrEqual
	LessThanOrEqual
	Equal
	NotEqual
	And
	Or
	Not
)

var tokenTypes = map[Type]string{
	Illegal:            "Illegal",
	Eof:                "Eof",
	Ident:              "Ident",
	Assign:             "Assign",
	Plus:               "Plus",
	Minus:              "Minus",
	Divide:             "Divide",
	Multiply:           "Multiply",
	Modulus:            "Modulus",
	Comma:              "Comma",
	FullStop:           "FullStop",
	Semicolon:          "Semicolon",
	Colon:              "Colon",
	LParen:             "LParen",
	RParen:             "RParen",
	LBrace:             "LBrace",
	RBrace:             "RBrace",
	LBracket:           "LBracket",
	RBracket:           "RBracket",
	Percent:            "Percent",
	DoubleQuote:        "DoubleQuote",
	SingleQuote:        "SingleQuote",
	GreaterThan:        "GreaterThan",
	LessThan:           "LessThan",
	GreaterThanOrEqual: "GreaterThanOrEqual",
	LessThanOrEqual:    "LessThanOrEqual",
	Equal:              "Equal",
	NotEqual:           "NotEqual",
	And:                "And",
	Or:                 "Or",
	Not:                "Not",
}

func GetTokenType[T any](t Token[T]) string {
	name, exists := tokenTypes[t.Type]

	if exists {
		return name
	}

	return "UNKNOWN"
}
