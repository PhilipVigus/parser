package token

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
