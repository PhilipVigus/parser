package token

type Type string
type Token struct {
	Type    Type
	Literal string
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
