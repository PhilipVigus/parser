package lexer

import (
	"os"
	"parser/lexer/token"
	"reflect"
	"strings"
	"testing"
)

func TestLexer_Tokenize_FromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			name:  "Single one character token",
			input: ";",
			expected: []token.Token{
				{
					Type:   token.Semicolon,
					Value:  ";",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 1,
				},
			},
		},
		{
			name:  "Greater than or equal token",
			input: ">=",
			expected: []token.Token{
				{
					Type:   token.GreaterThanOrEqual,
					Value:  ">=",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 2,
				},
			},
		},
		{
			name:  "Less than or equal token",
			input: "<=",
			expected: []token.Token{
				{
					Type:   token.LessThanOrEqual,
					Value:  "<=",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 2,
				},
			},
		},
		{
			name:  "Number token",
			input: "123",
			expected: []token.Token{
				{
					Type:   token.Number,
					Value:  "123",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 3,
				},
			},
		},
		{
			name:  "decimal number token",
			input: "0.123",
			expected: []token.Token{
				{
					Type:   token.Number,
					Value:  "0.123",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 5,
				},
			},
		},
		{
			name:  "Single identifier",
			input: "ident",
			expected: []token.Token{
				{
					Type:   token.Ident,
					Value:  "ident",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 5,
				},
			},
		},
		{
			name:  "Illegal character",
			input: "&",
			expected: []token.Token{
				{
					Type:   token.Illegal,
					Value:  "&",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 1,
				},
			},
		},
		{
			name:  "String with double quotes",
			input: "\"a string\"test",
			expected: []token.Token{
				{
					Type:   token.String,
					Value:  "a string",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Ident,
					Value:  "test",
					Line:   0,
					Column: 10,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 14,
				},
			},
		},
		{
			name:  "Non-terminating string with double quotes",
			input: "\"test",
			expected: []token.Token{
				{
					Type:   token.Illegal,
					Value:  "\"test",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 5,
				},
			},
		},
		{
			name:  "String with single quotes",
			input: "'a string'test",
			expected: []token.Token{
				{
					Type:   token.String,
					Value:  "a string",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Ident,
					Value:  "test",
					Line:   0,
					Column: 10,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 14,
				},
			},
		},
		{
			name:  "Non-terminating string with single quotes",
			input: "'a string++4",
			expected: []token.Token{
				{
					Type:   token.Illegal,
					Value:  "'a string++4",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 12,
				},
			},
		},
		{
			name:  "Multiple tokens",
			input: "ident_with_underscores;&/ident.",
			expected: []token.Token{
				{
					Type:   token.Ident,
					Value:  "ident_with_underscores",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Semicolon,
					Value:  ";",
					Line:   0,
					Column: 22,
				},
				{
					Type:   token.Illegal,
					Value:  "&",
					Line:   0,
					Column: 23,
				},
				{
					Type:   token.Divide,
					Value:  "/",
					Line:   0,
					Column: 24,
				},
				{
					Type:   token.Ident,
					Value:  "ident",
					Line:   0,
					Column: 25,
				},
				{
					Type:   token.FullStop,
					Value:  ".",
					Line:   0,
					Column: 30,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 31,
				},
			},
		},
		{
			name:  "Multiple tokens with whitespace",
			input: "ident_with_underscores; & / ident.",
			expected: []token.Token{
				{
					Type:   token.Ident,
					Value:  "ident_with_underscores",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Semicolon,
					Value:  ";",
					Line:   0,
					Column: 22,
				},
				{
					Type:   token.Illegal,
					Value:  "&",
					Line:   0,
					Column: 24,
				},
				{
					Type:   token.Divide,
					Value:  "/",
					Line:   0,
					Column: 26,
				},
				{
					Type:   token.Ident,
					Value:  "ident",
					Line:   0,
					Column: 28,
				},
				{
					Type:   token.FullStop,
					Value:  ".",
					Line:   0,
					Column: 33,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 34,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(strings.NewReader(tt.input))

			tokens := l.Tokenize()
			if !reflect.DeepEqual(tokens, tt.expected) {
				t.Errorf("Tokenize() = %v, want %v", tokens, tt.expected)
			}
		})
	}
}

func TestLexer_Tokenize_FromFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected []token.Token
	}{
		{
			name:     "Single line",
			filePath: "testdata/single_line.txt",
			expected: []token.Token{
				{
					Type:   token.Ident,
					Value:  "x",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Plus,
					Value:  "+",
					Line:   0,
					Column: 1,
				},
				{
					Type:   token.Number,
					Value:  "5",
					Line:   0,
					Column: 2,
				},
				{
					Type:   token.Semicolon,
					Value:  ";",
					Line:   0,
					Column: 3,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   0,
					Column: 4,
				},
			},
		},
		{
			name:     "Multiple lines",
			filePath: "testdata/multiple_lines.txt",
			expected: []token.Token{
				{
					Type:   token.Ident,
					Value:  "x",
					Line:   0,
					Column: 0,
				},
				{
					Type:   token.Assign,
					Value:  "=",
					Line:   0,
					Column: 2,
				},
				{
					Type:   token.Number,
					Value:  "5",
					Line:   0,
					Column: 4,
				},
				{
					Type:   token.Semicolon,
					Value:  ";",
					Line:   0,
					Column: 5,
				},
				{
					Type:   token.String,
					Value:  "test    test",
					Line:   3,
					Column: 0,
				},
				{
					Type:   token.Eof,
					Value:  "",
					Line:   3,
					Column: 14,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(tt.filePath)
			if err != nil {
				t.Errorf("Error opening test fixture: %v", err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					t.Errorf("Error closing test fixture: %v", err)
				}
			}(file)

			l := New(file)

			tokens := l.Tokenize()
			if !reflect.DeepEqual(tokens, tt.expected) {
				t.Errorf("Tokenize() = %v, want %v", tokens, tt.expected)
			}
		})
	}
}
