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
					Type:  token.Semicolon,
					Value: ";",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:  "Greater than or equal token",
			input: ">=",
			expected: []token.Token{
				{
					Type:  token.GreaterThanOrEqual,
					Value: ">=",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:  "Less than or equal token",
			input: "<=",
			expected: []token.Token{
				{
					Type:  token.LessThanOrEqual,
					Value: "<=",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:  "Number token",
			input: "123",
			expected: []token.Token{
				{
					Type:  token.Number,
					Value: "123",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:  "decimal number token",
			input: "0.123",
			expected: []token.Token{
				{
					Type:  token.Number,
					Value: "0.123",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:  "Single identifier",
			input: "ident",
			expected: []token.Token{
				{
					Type:  token.Ident,
					Value: "ident",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:  "Illegal character",
			input: "&",
			expected: []token.Token{
				{
					Type:  token.Illegal,
					Value: "&",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:  "String with double quotes",
			input: "\"a string\"test",
			expected: []token.Token{
				{
					Type:  token.String,
					Value: "a string",
				},
				{
					Type:  token.Ident,
					Value: "test",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:  "Non-terminating string with double quotes",
			input: "\"test",
			expected: []token.Token{
				{
					Type:  token.Illegal,
					Value: "\"test",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:  "String with single quotes",
			input: "'a string'test",
			expected: []token.Token{
				{
					Type:  token.String,
					Value: "a string",
				},
				{
					Type:  token.Ident,
					Value: "test",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:  "Non-terminating string with single quotes",
			input: "'a string++4",
			expected: []token.Token{
				{
					Type:  token.Illegal,
					Value: "'a string++4",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:  "Multiple tokens",
			input: "ident_with_underscores;&/ident.",
			expected: []token.Token{
				{
					Type:  token.Ident,
					Value: "ident_with_underscores",
				},
				{
					Type:  token.Semicolon,
					Value: ";",
				},
				{
					Type:  token.Illegal,
					Value: "&",
				},
				{
					Type:  token.Divide,
					Value: "/",
				},
				{
					Type:  token.Ident,
					Value: "ident",
				},
				{
					Type:  token.FullStop,
					Value: ".",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:  "Multiple tokens with whitespace",
			input: "ident_with_underscores; & / ident.",
			expected: []token.Token{
				{
					Type:  token.Ident,
					Value: "ident_with_underscores",
				},
				{
					Type:  token.Semicolon,
					Value: ";",
				},
				{
					Type:  token.Illegal,
					Value: "&",
				},
				{
					Type:  token.Divide,
					Value: "/",
				},
				{
					Type:  token.Ident,
					Value: "ident",
				},
				{
					Type:  token.FullStop,
					Value: ".",
				},
				{
					Type:  token.Eof,
					Value: "",
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
					Type:  token.Ident,
					Value: "x",
				},
				{
					Type:  token.Plus,
					Value: "+",
				},
				{
					Type:  token.Number,
					Value: "5",
				},
				{
					Type:  token.Semicolon,
					Value: ";",
				},
				{
					Type:  token.Eof,
					Value: "",
				},
			},
		},
		{
			name:     "Multiple lines",
			filePath: "testdata/multiple_lines.txt",
			expected: []token.Token{
				{
					Type:  token.Ident,
					Value: "x",
				},
				{
					Type:  token.Assign,
					Value: "=",
				},
				{
					Type:  token.Number,
					Value: "5",
				},
				{
					Type:  token.Semicolon,
					Value: ";",
				},
				{
					Type:  token.String,
					Value: "test    test",
				},
				{
					Type:  token.Eof,
					Value: "",
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
