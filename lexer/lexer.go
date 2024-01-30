package lexer

import (
	"fmt"
	"parser/lexer/token"
	"unicode/utf8"
)

type Lexer struct {
	position     int
	readPosition int
	input        string
	inputLength  int
	ch           rune
}

func New(s string) (*Lexer, error) {
	if s == "" {
		return nil, fmt.Errorf("empty input string")
	}
	l := &Lexer{
		input:       s,
		inputLength: utf8.RuneCountInString(s),
	}
	l.readNextChar()
	return l, nil
}

func (l *Lexer) readNextChar() {
	if l.readPosition >= l.inputLength {
		l.position = l.readPosition
		l.readPosition++

		l.ch = 0
		return
	}

	c, size := utf8.DecodeRuneInString(l.input[l.readPosition:])

	if c == utf8.RuneError {
		if size == 0 {
			l.ch = 0
		} else {
			l.ch = utf8.RuneError
		}
	} else {
		l.ch = c
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token[any] {
	var t token.Token[any]

	switch l.ch {
	case '+':
		t = token.New(token.PLUS, "+")
		l.readNextChar()
	case '-':
		t = token.New(token.MINUS, "-")
		l.readNextChar()
	case 0:
		t = token.New(token.EOF, "")
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			t = token.New(token.IDENT, ident)
		} else {
			t = token.New(token.ILLEGAL, string(l.ch))
			l.readNextChar()
		}
	}

	return t
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position // Mark the start of the identifier
	for isLetter(l.ch) {
		l.readNextChar()
	}
	// No need to adjust l.position or l.readPosition here
	return l.input[startPosition:l.position]
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
