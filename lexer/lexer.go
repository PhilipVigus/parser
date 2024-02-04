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

func (l *Lexer) peekNextChar() rune {
	if l.readPosition >= l.inputLength {
		return 0
	}
	c, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
	return c
}

func (l *Lexer) Tokenize() []token.Token[any] {
	var tokens []token.Token[any]
	for {
		t := l.NextToken()
		tokens = append(tokens, t)
		if t.Type == token.Eof {
			break
		}
	}
	return tokens
}

func (l *Lexer) NextToken() token.Token[any] {
	var t token.Token[any]

	switch l.ch {
	case '=':
		if l.peekNextChar() == '=' {
			l.readNextChar()
			t = token.New(token.Equal, "==")
		} else {
			t = token.New(token.Assign, "=")
		}
		l.readNextChar()
	case '+':
		t = token.New(token.Plus, "+")
		l.readNextChar()
	case '-':
		t = token.New(token.Minus, "-")
		l.readNextChar()
	case '*':
		t = token.New(token.Multiply, "*")
		l.readNextChar()
	case '/':
		t = token.New(token.Divide, "/")
		l.readNextChar()
	case ',':
		t = token.New(token.Comma, ",")
		l.readNextChar()
	case '.':
		t = token.New(token.FullStop, ".")
		l.readNextChar()
	case ';':
		t = token.New(token.Semicolon, ";")
		l.readNextChar()
	case ':':
		t = token.New(token.Colon, ":")
		l.readNextChar()
	case '(':
		t = token.New(token.LParen, "(")
		l.readNextChar()
	case ')':
		t = token.New(token.RParen, ")")
		l.readNextChar()
	case '{':
		t = token.New(token.LBrace, "{")
		l.readNextChar()
	case '}':
		t = token.New(token.RBrace, "}")
		l.readNextChar()
	case '[':
		t = token.New(token.LBracket, "[")
		l.readNextChar()
	case ']':
		t = token.New(token.RBracket, "]")
		l.readNextChar()
	case '%':
		t = token.New(token.Percent, "%")
		l.readNextChar()
	case '"':
		t = token.New(token.DoubleQuote, "\"")
		l.readNextChar()
	case '\'':
		t = token.New(token.SingleQuote, "'")
		l.readNextChar()
	case '>':
		if l.peekNextChar() == '=' {
			l.readNextChar()
			t = token.New(token.GreaterThanOrEqual, ">=")
		} else {
			t = token.New(token.GreaterThan, ">")
		}
		l.readNextChar()
	case '<':
		if l.peekNextChar() == '=' {
			l.readNextChar()
			t = token.New(token.LessThanOrEqual, "<=")
		} else {
			t = token.New(token.LessThan, "<")
		}
		l.readNextChar()
	case '!':
		if l.peekNextChar() == '=' {
			l.readNextChar()
			t = token.New(token.NotEqual, "!=")
		} else {
			t = token.New(token.Not, "!")
		}
		l.readNextChar()
	case 0:
		t = token.New(token.Eof, "")
	default:
		if isLetter(l.ch) {
			Ident := l.readIdentifier()
			t = token.New(token.Ident, Ident)
		} else {
			t = token.New(token.Illegal, string(l.ch))
			l.readNextChar()
		}
	}

	return t
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position // Mark the start of the Identifier
	for isLetter(l.ch) {
		l.readNextChar()
	}
	// No need to adjust l.position or l.readPosition here
	return l.input[startPosition:l.position]
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
