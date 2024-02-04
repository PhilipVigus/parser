package token

import (
	"reflect"
	"testing"
)

func TestToken_New(t *testing.T) {
	tests := []struct {
		name     string
		t        Type
		lit      any
		expected Token[any]
	}{
		{
			name: "Illegal",
			t:    Illegal,
			lit:  "Illegal",
			expected: Token[any]{
				Type:    Illegal,
				Literal: "Illegal",
			},
		},
		{
			name: "Eof",
			t:    Eof,
			lit:  "Eof",
			expected: Token[any]{
				Type:    Eof,
				Literal: "Eof",
			},
		},
		{
			name: "Ident",
			t:    Ident,
			lit:  "Ident",
			expected: Token[any]{
				Type:    Ident,
				Literal: "Ident",
			},
		},
		{
			name: "Assign",
			t:    Assign,
			lit:  "Assign",
			expected: Token[any]{
				Type:    Assign,
				Literal: "Assign",
			},
		},
		{
			name: "Plus",
			t:    Plus,
			lit:  "Plus",
			expected: Token[any]{
				Type:    Plus,
				Literal: "Plus",
			},
		},
		{
			name: "Minus",
			t:    Minus,
			lit:  "Minus",
			expected: Token[any]{
				Type:    Minus,
				Literal: "Minus",
			},
		},
		{
			name: "Divide",
			t:    Divide,
			lit:  "Divide",
			expected: Token[any]{
				Type:    Divide,
				Literal: "Divide",
			},
		},
		{
			name: "Multiply",
			t:    Multiply,
			lit:  "Multiply",
			expected: Token[any]{
				Type:    Multiply,
				Literal: "Multiply",
			},
		},
		{
			name: "Modulus",
			t:    Modulus,
			lit:  "Modulus",
			expected: Token[any]{
				Type:    Modulus,
				Literal: "Modulus",
			},
		},
		{
			name: "Comma",
			t:    Comma,
			lit:  "Comma",
			expected: Token[any]{
				Type:    Comma,
				Literal: "Comma",
			},
		},
		{
			name: "FullStop",
			t:    FullStop,
			lit:  "FullStop",
			expected: Token[any]{
				Type:    FullStop,
				Literal: "FullStop",
			},
		},
		{
			name: "Semicolon",
			t:    Semicolon,
			lit:  "Semicolon",
			expected: Token[any]{
				Type:    Semicolon,
				Literal: "Semicolon",
			},
		},
		{
			name: "Colon",
			t:    Colon,
			lit:  "Colon",
			expected: Token[any]{
				Type:    Colon,
				Literal: "Colon",
			},
		},
		{
			name: "LParen",
			t:    LParen,
			lit:  "LParen",
			expected: Token[any]{
				Type:    LParen,
				Literal: "LParen",
			},
		},
		{
			name: "RParen",
			t:    RParen,
			lit:  "RParen",
			expected: Token[any]{
				Type:    RParen,
				Literal: "RParen",
			},
		},
		{
			name: "LBrace",
			t:    LBrace,
			lit:  "LBrace",
			expected: Token[any]{
				Type:    LBrace,
				Literal: "LBrace",
			},
		},
		{
			name: "RBrace",
			t:    RBrace,
			lit:  "RBrace",
			expected: Token[any]{
				Type:    RBrace,
				Literal: "RBrace",
			},
		},
		{
			name: "LBracket",
			t:    LBracket,
			lit:  "LBracket",
			expected: Token[any]{
				Type:    LBracket,
				Literal: "LBracket",
			},
		},
		{
			name: "RBracket",
			t:    RBracket,
			lit:  "RBracket",
			expected: Token[any]{
				Type:    RBracket,
				Literal: "RBracket",
			},
		},
		{
			name: "Percent",
			t:    Percent,
			lit:  "Percent",
			expected: Token[any]{
				Type:    Percent,
				Literal: "Percent",
			},
		},
		{
			name: "DoubleQuote",
			t:    DoubleQuote,
			lit:  "DoubleQuote",
			expected: Token[any]{
				Type:    DoubleQuote,
				Literal: "DoubleQuote",
			},
		},
		{
			name: "SingleQuote",
			t:    SingleQuote,
			lit:  "SingleQuote",
			expected: Token[any]{
				Type:    SingleQuote,
				Literal: "SingleQuote",
			},
		},
		{
			name: "GreaterThan",
			t:    GreaterThan,
			lit:  "GreaterThan",
			expected: Token[any]{
				Type:    GreaterThan,
				Literal: "GreaterThan",
			},
		},
		{
			name: "LessThan",
			t:    LessThan,
			lit:  "LessThan",
			expected: Token[any]{
				Type:    LessThan,
				Literal: "LessThan",
			},
		},
		{
			name: "GreaterThanOrEqual",
			t:    GreaterThanOrEqual,
			lit:  "GreaterThanOrEqual",
			expected: Token[any]{
				Type:    GreaterThanOrEqual,
				Literal: "GreaterThanOrEqual",
			},
		},
		{
			name: "LessThanOrEqual",
			t:    LessThanOrEqual,
			lit:  "LessThanOrEqual",
			expected: Token[any]{
				Type:    LessThanOrEqual,
				Literal: "LessThanOrEqual",
			},
		},
		{
			name: "Equal",
			t:    Equal,
			lit:  "Equal",
			expected: Token[any]{
				Type:    Equal,
				Literal: "Equal",
			},
		},
		{
			name: "NotEqual",
			t:    NotEqual,
			lit:  "NotEqual",
			expected: Token[any]{
				Type:    NotEqual,
				Literal: "NotEqual",
			},
		},
		{
			name: "And",
			t:    And,
			lit:  "And",
			expected: Token[any]{
				Type:    And,
				Literal: "And",
			},
		},
		{
			name: "Or",
			t:    Or,
			lit:  "Or",
			expected: Token[any]{
				Type:    Or,
				Literal: "Or",
			},
		},
		{
			name: "Not",
			t:    Not,
			lit:  "Not",
			expected: Token[any]{
				Type:    Not,
				Literal: "Not",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := New(tt.t, tt.lit)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected: %v, Actual: %v", tt.expected, actual)
			}
		})
	}
}

func TestToken_String(t *testing.T) {
	tests := []struct {
		name     string
		t        Type
		lit      any
		expected string
	}{
		{
			name:     "Illegal",
			t:        Illegal,
			lit:      "Illegal",
			expected: "[ type: Illegal, literal: Illegal ]",
		},
		{
			name:     "Eof",
			t:        Eof,
			lit:      "Eof",
			expected: "[ type: Eof, literal: Eof ]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := New(tt.t, tt.lit).String()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected: %v, Actual: %v", tt.expected, actual)
			}
		})
	}
}

func TestGetTokenType(t *testing.T) {
	tests := []struct {
		name     string
		t        Token[any]
		expected string
	}{
		{
			name: "Illegal",
			t: Token[any]{
				Type:    Illegal,
				Literal: "Illegal",
			},
			expected: "Illegal",
		},
		{
			name: "Eof",
			t: Token[any]{
				Type:    Eof,
				Literal: "Eof",
			},
			expected: "Eof",
		},
		{
			name: "Ident",
			t: Token[any]{
				Type:    Ident,
				Literal: "Ident",
			},
			expected: "Ident",
		},
		{
			name: "Assign",
			t: Token[any]{
				Type:    Assign,
				Literal: "Assign",
			},
			expected: "Assign",
		},
		{
			name: "Plus",
			t: Token[any]{
				Type:    Plus,
				Literal: "Plus",
			},
			expected: "Plus",
		},
		{
			name: "Minus",
			t: Token[any]{
				Type:    Minus,
				Literal: "Minus",
			},
			expected: "Minus",
		},
		{
			name: "Divide",
			t: Token[any]{
				Type:    Divide,
				Literal: "Divide",
			},
			expected: "Divide",
		},
		{
			name: "Multiply",
			t: Token[any]{
				Type:    Multiply,
				Literal: "Multiply",
			},
			expected: "Multiply",
		},
		{
			name: "Comma",
			t: Token[any]{
				Type:    Comma,
				Literal: "Comma",
			},
			expected: "Comma",
		},
		{
			name: "FullStop",
			t: Token[any]{
				Type:    FullStop,
				Literal: "FullStop",
			},
			expected: "FullStop",
		},
		{
			name: "Semicolon",
			t: Token[any]{
				Type:    Semicolon,
				Literal: "Semicolon",
			},
			expected: "Semicolon",
		},
		{
			name: "Colon",
			t: Token[any]{
				Type:    Colon,
				Literal: "Colon",
			},
			expected: "Colon",
		},
		{
			name: "LParen",
			t: Token[any]{
				Type:    LParen,
				Literal: "LParen",
			},
			expected: "LParen",
		},
		{
			name: "RParen",
			t: Token[any]{
				Type:    RParen,
				Literal: "RParen",
			},
			expected: "RParen",
		},
		{
			name: "LBrace",
			t: Token[any]{
				Type:    LBrace,
				Literal: "LBrace",
			},
			expected: "LBrace",
		},
		{
			name: "RBrace",
			t: Token[any]{
				Type:    RBrace,
				Literal: "RBrace",
			},
			expected: "RBrace",
		},
		{
			name: "LBracket",
			t: Token[any]{
				Type:    LBracket,
				Literal: "LBracket",
			},
			expected: "LBracket",
		},
		{
			name: "RBracket",
			t: Token[any]{
				Type:    RBracket,
				Literal: "RBracket",
			},
			expected: "RBracket",
		},
		{
			name: "Percent",
			t: Token[any]{
				Type:    Percent,
				Literal: "Percent",
			},
			expected: "Percent",
		},
		{
			name: "DoubleQuote",
			t: Token[any]{
				Type:    DoubleQuote,
				Literal: "DoubleQuote",
			},
			expected: "DoubleQuote",
		},
		{
			name: "SingleQuote",
			t: Token[any]{
				Type:    SingleQuote,
				Literal: "SingleQuote",
			},
			expected: "SingleQuote",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GetTokenType(tt.t)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected: %v, Actual: %v", tt.expected, actual)
			}
		})
	}

}
