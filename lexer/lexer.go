package lexer

import (
	"bufio"
	"errors"
	"io"
	"parser/lexer/token"
	"strings"
	"unicode"
)

type Lexer struct {
	reader *bufio.Reader

	position     int
	readPosition int
	input        string
	inputLength  int
	ch           rune
	width        int
}

func New(r io.Reader) *Lexer {
	l := &Lexer{
		reader: bufio.NewReader(r),
	}
	l.readNextChar()
	return l
}

func (l *Lexer) readNextChar() {
	var err error

	l.ch, l.width, err = l.reader.ReadRune()

	if err == nil {
		return
	}

	if err == io.EOF {
		l.ch = 0
		l.width = 0
	} else {
		panic(err)
	}
}

func (l *Lexer) peekNextChar() rune {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			return 0
		}
		panic(err)
	}

	err = l.reader.UnreadRune()
	if err != nil {
		panic(err)
	}
	return r
}

func (l *Lexer) Tokenize() []token.Token {
	var tokens []token.Token
	for {
		t := l.NextToken()
		tokens = append(tokens, t)
		if t.Type == token.Eof {
			break
		}
	}
	return tokens
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

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
	case '&':
		if l.peekNextChar() == '&' {
			l.readNextChar()
			t = token.New(token.And, "&&")
		} else {
			t = token.New(token.Illegal, string(l.ch))
		}
		l.readNextChar()
	case '|':
		if l.peekNextChar() == '|' {
			l.readNextChar()
			t = token.New(token.Or, "||")
		} else {
			t = token.New(token.Illegal, string(l.ch))
		}
		l.readNextChar()
	case '"':
		str, err := l.readString('"')
		if err {
			t = token.New(token.Illegal, str)
		} else {
			t = token.New(token.String, str)
		}
		l.readNextChar()
	case '\'':
		str, err := l.readString('\'')
		if err {
			t = token.New(token.Illegal, str)
		} else {
			t = token.New(token.String, str)
		}
		l.readNextChar()
	case 0:
		t = token.New(token.Eof, "")
	default:
		if isLetter(l.ch) {
			t = l.readIdentifier()
		} else if isDigit(l.ch) {
			num, err := l.readNumber()

			if err == nil {
				t = token.New(token.Number, num)
			} else {
				t = token.New(token.Illegal, num)
			}
		} else {
			t = token.New(token.Illegal, string(l.ch))
			l.readNextChar()
		}
	}

	return t
}

func (l *Lexer) readIdentifier() token.Token {
	var builder strings.Builder
	for unicode.IsLetter(l.ch) || l.ch == '_' {
		builder.WriteRune(l.ch)
		l.readNextChar()
	}
	identifier := builder.String()

	return token.New(token.GetKeywordType(identifier), identifier)
}

func (l *Lexer) readNumber() (string, error) {
	var builder strings.Builder
	haveReadDot := false

	for unicode.IsDigit(l.ch) || (l.ch == '.' && !haveReadDot && unicode.IsDigit(l.peekNextChar())) {
		if l.ch == '.' {
			if haveReadDot {
				break
			}
			haveReadDot = true
		}
		builder.WriteRune(l.ch)
		l.readNextChar()
	}

	str := builder.String()

	if l.endsWithDot(str, haveReadDot) {
		return str, errors.New("number ends with a dot")
	}
	return str, nil
}

func (l *Lexer) endsWithDot(str string, haveReadDot bool) bool {
	notEmpty := len(str) > 0
	lastCharIsDot := str[len(str)-1] == '.'
	endsWithDot := haveReadDot && notEmpty && lastCharIsDot
	return endsWithDot
}

func (l *Lexer) readString(ch rune) (string, bool) {
	var builder strings.Builder

	initialQuote := l.ch

	l.readNextChar()

	for {
		if l.ch == ch || l.ch == 0 {
			break
		}
		builder.WriteRune(l.ch)
		l.readNextChar()
	}

	s := builder.String()

	if l.ch == 0 {
		s = string(initialQuote) + builder.String()
	}

	return s, l.ch == 0
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}
