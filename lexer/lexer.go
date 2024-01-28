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
	ch           rune
}

func New(s string) (*Lexer, error) {
	if s == "" {
		return nil, fmt.Errorf("empty input string")
	}
	l := &Lexer{
		input: s,
	}
	l.readNextChar()
	return l, nil
}

func (l *Lexer) readNextChar() (rune, int) {
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
	return c, size
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	switch l.ch {
	case '+':
		t = token.New(token.PLUS, '+')
	case '-':
		t = token.New(token.MINUS, '-')
	case 0:
		t = token.New(token.EOF, 0)
	default:
		t = token.New(token.ILLEGAL, 0)
	}

	if t.Type != token.EOF && t.Type != token.ILLEGAL {
		l.readNextChar()
	}

	return t
}
