package lexer

import (
	"parser/lexer/token"
	"reflect"
	"testing"
)

func TestLexer_Tokenize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.Token[any]
	}{
		{
			name:  "Single one character token",
			input: ";",
			expected: []token.Token[any]{
				{
					Type:    token.Semicolon,
					Literal: ";",
				},
				{
					Type:    token.Eof,
					Literal: "",
				},
			},
		},
		{
			name:  "Greater than or equal token",
			input: ">=",
			expected: []token.Token[any]{
				{
					Type:    token.GreaterThanOrEqual,
					Literal: ">=",
				},
				{
					Type:    token.Eof,
					Literal: "",
				},
			},
		},
		{
			name:  "Less than or equal token",
			input: "<=",
			expected: []token.Token[any]{
				{
					Type:    token.LessThanOrEqual,
					Literal: "<=",
				},
				{
					Type:    token.Eof,
					Literal: "",
				},
			},
		},
		{
			name:  "Number token",
			input: "123",
			expected: []token.Token[any]{
				{
					Type:    token.Number,
					Literal: "123",
				},
				{
					Type:    token.Eof,
					Literal: "",
				},
			},
		},
		{
			name:  "decimal number token",
			input: "0.123",
			expected: []token.Token[any]{
				{
					Type:    token.Number,
					Literal: "0.123",
				},
				{
					Type:    token.Eof,
					Literal: "",
				},
			},
		},
		{
			name:  "Single identifier",
			input: "ident",
			expected: []token.Token[any]{
				{
					Type:    token.Ident,
					Literal: "ident",
				},
				{
					Type:    token.Eof,
					Literal: "",
				},
			},
		},
		{
			name:  "Illegal character",
			input: "&",
			expected: []token.Token[any]{
				{
					Type:    token.Illegal,
					Literal: "&",
				},
				{
					Type:    token.Eof,
					Literal: "",
				},
			},
		},
		{
			name:  "String with double quotes",
			input: "\"a string\"test",
			expected: []token.Token[any]{
				{
					Type:    token.String,
					Literal: "a string",
				},
				{
					Type:    token.Ident,
					Literal: "test",
				},
				{
					Type:    token.Eof,
					Literal: "",
				},
			},
		},
		{
			name:  "Non-terminating string with double quotes",
			input: "\"a string++4",
			expected: []token.Token[any]{
				{
					Type:    token.Illegal,
					Literal: "\"a string++4",
				},
				{
					Type:    token.Eof,
					Literal: "",
				},
			},
		},
		{
			name:  "String with single quotes",
			input: "'a string'test",
			expected: []token.Token[any]{
				{
					Type:    token.String,
					Literal: "a string",
				},
				{
					Type:    token.Ident,
					Literal: "test",
				},
				{
					Type:    token.Eof,
					Literal: "",
				},
			},
		},
		{
			name:  "Non-terminating string with single quotes",
			input: "'a string++4",
			expected: []token.Token[any]{
				{
					Type:    token.Illegal,
					Literal: "'a string++4",
				},
				{
					Type:    token.Eof,
					Literal: "",
				},
			},
		},
		{
			name:  "Multiple tokens",
			input: "ident_with_underscores;&/ident.",
			expected: []token.Token[any]{
				{
					Type:    token.Ident,
					Literal: "ident_with_underscores",
				},
				{
					Type:    token.Semicolon,
					Literal: ";",
				},
				{
					Type:    token.Illegal,
					Literal: "&",
				},
				{
					Type:    token.Divide,
					Literal: "/",
				},
				{
					Type:    token.Ident,
					Literal: "ident",
				},
				{
					Type:    token.FullStop,
					Literal: ".",
				},
				{
					Type:    token.Eof,
					Literal: "",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := New(tt.input)
			if err != nil {
				t.Fatalf("Lexer creation failed: %v", err)
			}

			tokens := l.Tokenize()
			if !reflect.DeepEqual(tokens, tt.expected) {
				t.Errorf("Tokenize() = %v, want %v", tokens, tt.expected)
			}
		})
	}
}
