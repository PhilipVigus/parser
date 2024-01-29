package token

import "fmt"

type Type int
type Token struct {
	Type    Type
	Literal rune
}

func New(t Type, l rune) Token {
	return Token{
		Type:    t,
		Literal: l,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("[ type: %s, literal: %c ]", GetTokenType(t), t.Literal)
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

func GetTokenType(t Token) string {
	name, exists := tokenTypes[t.Type]

	if exists {
		return name
	}
	return "UNKNOWN"
}
