// Package lexer provides a lexer for the  programming language.
package lexer

import (
	"bufio"
	"errors"
	"io"
	"parser/lexer/token"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Lexer is a lexer for the programming language.
type Lexer struct {
	// reader is the buffer containing the input string.
	reader *bufio.Reader
	// ch is the current character being read.
	ch rune
	// line is the current line number.
	line int
	// col is the current column number.
	col int
}

// New creates a new lexer from the given reader.
func New(r io.Reader) *Lexer {
	l := &Lexer{
		reader: bufio.NewReader(r),
		line:   0,
		col:    0,
	}
	// Read the first character to initialize the lexer.
	l.readNextChar()
	return l
}

// readNextChar reads the next character from the input string.
func (l *Lexer) readNextChar() {
	var err error
	l.ch, _, err = l.reader.ReadRune()
	l.col++
	if err == nil {
		return
	}

	if err == io.EOF {
		l.ch = 0
	} else {
		panic(err)
	}
}

// peekNextChar returns the next character from the input string without consuming it.
func (l *Lexer) peekNextChar() (rune, error) {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			return 0, nil
		}
		return 0, err
	}

	// 'Unread' the rune so that it can be read again.
	if err := l.reader.UnreadRune(); err != nil {
		return 0, err
	}

	return r, nil
}

// Tokenize reads the entire input string and returns a slice of tokens.
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

// NextToken returns the next token from the input string.
func (l *Lexer) NextToken() token.Token {
	var t token.Token

	if isWhitespace(l.ch) {
		return l.handleWhitespace()
	}

	switch l.ch {
	case '=':
		nextChar, err := l.peekNextChar()

		if err != nil {
			t = token.New(token.Illegal, string(l.ch), l.line, l.col-1)
			l.readNextChar()
			break
		}

		if nextChar == '=' {
			l.readNextChar()
			t = token.New(token.Equal, "==", l.line, l.col-2)
		} else {
			t = token.New(token.Assign, "=", l.line, l.col-1)
		}
		l.readNextChar()
	case '+':
		t = token.New(token.Plus, "+", l.line, l.col-1)
		l.readNextChar()
	case '-':
		t = token.New(token.Minus, "-", l.line, l.col-1)
		l.readNextChar()
	case '*':
		t = token.New(token.Multiply, "*", l.line, l.col-1)
		l.readNextChar()
	case '/':
		t = token.New(token.Divide, "/", l.line, l.col-1)
		l.readNextChar()
	case ',':
		t = token.New(token.Comma, ",", l.line, l.col-1)
		l.readNextChar()
	case '.':
		t = token.New(token.FullStop, ".", l.line, l.col-1)
		l.readNextChar()
	case ';':
		t = token.New(token.Semicolon, ";", l.line, l.col-1)
		l.readNextChar()
	case ':':
		t = token.New(token.Colon, ":", l.line, l.col-1)
		l.readNextChar()
	case '(':
		t = token.New(token.LParen, "(", l.line, l.col-1)
		l.readNextChar()
	case ')':
		t = token.New(token.RParen, ")", l.line, l.col-1)
		l.readNextChar()
	case '{':
		t = token.New(token.LBrace, "{", l.line, l.col-1)
		l.readNextChar()
	case '}':
		t = token.New(token.RBrace, "}", l.line, l.col-1)
		l.readNextChar()
	case '[':
		t = token.New(token.LBracket, "[", l.line, l.col-1)
		l.readNextChar()
	case ']':
		t = token.New(token.RBracket, "]", l.line, l.col-1)
		l.readNextChar()
	case '%':
		t = token.New(token.Percent, "%", l.line, l.col-1)
		l.readNextChar()
	case '>':
		nextChar, err := l.peekNextChar()

		if err != nil {
			t = token.New(token.Illegal, string(l.ch), l.line, l.col-1)
			l.readNextChar()
			break
		}

		if nextChar == '=' {
			l.readNextChar()
			t = token.New(token.GreaterThanOrEqual, ">=", l.line, l.col-2)
		} else {
			t = token.New(token.GreaterThan, ">", l.line, l.col-1)
		}
		l.readNextChar()
	case '<':
		nextChar, err := l.peekNextChar()

		if err != nil {
			t = token.New(token.Illegal, string(l.ch), l.line, l.col-1)
			l.readNextChar()
			break
		}

		if nextChar == '=' {
			l.readNextChar()
			t = token.New(token.LessThanOrEqual, "<=", l.line, l.col-2)
		} else {
			t = token.New(token.LessThan, "<", l.line, l.col-1)
		}
		l.readNextChar()
	case '!':
		nextChar, err := l.peekNextChar()

		if err != nil {
			t = token.New(token.Illegal, string(l.ch), l.line, l.col-1)
			l.readNextChar()
			break
		}

		if nextChar == '=' {
			l.readNextChar()
			t = token.New(token.NotEqual, "!=", l.line, l.col-2)
		} else {
			t = token.New(token.Not, "!", l.line, l.col-1)
		}
		l.readNextChar()
	case '&':
		nextChar, err := l.peekNextChar()

		if err != nil {
			t = token.New(token.Illegal, string(l.ch), l.line, l.col-1)
			l.readNextChar()
			break
		}

		if nextChar == '&' {
			l.readNextChar()
			t = token.New(token.And, "&&", l.line, l.col-2)
		} else {
			t = token.New(token.Illegal, string(l.ch), l.line, l.col-1)
		}
		l.readNextChar()
	case '|':
		nextChar, err := l.peekNextChar()

		if err != nil {
			t = token.New(token.Illegal, string(l.ch), l.line, l.col-1)
			l.readNextChar()
			break
		}

		if nextChar == '|' {
			l.readNextChar()
			t = token.New(token.Or, "||", l.line, l.col-2)
		} else {
			t = token.New(token.Illegal, string(l.ch), l.line, l.col-1)
		}
		l.readNextChar()
	case '"':
		str, err := l.readString('"')
		if err {
			t = token.New(token.Illegal, str, l.line, l.col-utf8.RuneCountInString(str)-1)
			l.col--
		} else {
			t = token.New(token.String, str, l.line, l.col-utf8.RuneCountInString(str)-2)
		}
		l.readNextChar()
	case '\'':
		str, err := l.readString('\'')
		if err {
			t = token.New(token.Illegal, str, l.line, l.col-utf8.RuneCountInString(str)-1)
			l.col--
		} else {
			t = token.New(token.String, str, l.line, l.col-utf8.RuneCountInString(str)-2)
		}
		l.readNextChar()
	case 0:
		t = l.handleEof()
	default:
		t = l.handleDefaultCase(t)
	}

	return t
}

func (l *Lexer) handleEof() token.Token {
	return token.New(token.Eof, "", l.line, l.col-1)
}

func (l *Lexer) handleDefaultCase(t token.Token) token.Token {
	switch {
	case isLetter(l.ch):
		t = l.handleIdentifier(t)
	case unicode.IsDigit(l.ch):
		t = l.handleNumber(t)
	default:
		t = l.handleIllegalRune(t)
	}
	return t
}

func (l *Lexer) handleIllegalRune(t token.Token) token.Token {
	t = token.New(token.Illegal, string(l.ch), l.line, l.col-1)
	l.readNextChar()
	return t
}

func (l *Lexer) handleNumber(t token.Token) token.Token {
	num, err := l.readNumber()

	if err == nil {
		t = token.New(token.Number, num, l.line, l.col-utf8.RuneCountInString(num)-1)
	} else {
		t = token.New(token.Illegal, num, l.line, l.col-utf8.RuneCountInString(num)-1)
	}
	return t
}

func (l *Lexer) handleIdentifier(t token.Token) token.Token {
	t = l.readIdentifier()
	return t
}

func (l *Lexer) handleWhitespace() token.Token {
	if l.ch == '\n' {
		l.line++
		l.col = 0
	}
	l.readNextChar()
	return l.NextToken()
}

func (l *Lexer) readIdentifier() token.Token {
	var builder strings.Builder
	for unicode.IsLetter(l.ch) || l.ch == '_' {
		builder.WriteRune(l.ch)
		l.readNextChar()
	}
	identifier := builder.String()

	return token.New(token.GetKeywordType(identifier), identifier, l.line, l.col-utf8.RuneCountInString(identifier)-1)
}

func (l *Lexer) readNumber() (string, error) {
	var builder strings.Builder
	haveReadDot := false

	for {
		// Check if the current character is not a digit and not a valid dot condition.
		if !unicode.IsDigit(l.ch) && (l.ch != '.' || haveReadDot) {
			break
		}

		// Handle dot followed by a digit
		if l.ch == '.' {
			nextChar, err := l.peekNextChar()
			if err != nil {
				if errors.Is(err, io.EOF) {
					// If EOF, it's a valid end of number, break the loop
					break
				}
				// Return error immediately if not EOF
				return "", err
			}

			if !unicode.IsDigit(nextChar) {
				break
			}

			haveReadDot = true
		}

		builder.WriteRune(l.ch)
		l.readNextChar()
	}

	str := builder.String()

	// Check for invalid number format
	if l.numEndsWithDot(str, haveReadDot) {
		return str, errors.New("number ends with a dot")
	}

	// Return the number string and no error
	return str, nil
}

func (l *Lexer) numEndsWithDot(str string, haveReadDot bool) bool {
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

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
