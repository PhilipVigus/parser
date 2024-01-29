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

func (l *Lexer) peekNextChar() (rune, int) {
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

	return c, size
}

func (l *Lexer) NextToken() token.Token[any] {
	var t token.Token[any]

	switch l.ch {
	case '+':
		t = token.New(token.PLUS, "+")
	case '-':
		t = token.New(token.MINUS, "-")
	case 0:
		t = token.New(token.EOF, 0)
	default:
		if isLetter(l.ch) {
			t = token.New(token.IDENT, l.readIdentifier())
		} else {
			t = token.New(token.ILLEGAL, l.ch)
		}
	}

	if t.Type != token.EOF && t.Type != token.ILLEGAL {
		l.readNextChar()
	}

	return t
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readNextChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
