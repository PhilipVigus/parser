package lexer

import (
	"unicode/utf8"
)

type Lexer struct {
	charIndex int
	str       string
}

func NewLexer(s string) Lexer {
	return Lexer{
		charIndex: 0,
		str:       s,
	}
}

func (l *Lexer) GetNextChar() (rune, int) {
	c, size := utf8.DecodeRuneInString(l.str[l.charIndex:])
	l.charIndex++
	return c, size
}

func (l *Lexer) PeekNextChar() (rune, int) {
	c, size := utf8.DecodeRuneInString(l.str[l.charIndex:])
	return c, size
}
