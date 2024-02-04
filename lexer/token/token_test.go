package token

import (
	"reflect"
	"testing"
)

func TestToken_New(t *testing.T) {
	tests := []struct {
		name     string
		t        Type
		lit      string
		expected Token
	}{
		{
			name: "Illegal",
			t:    Illegal,
			lit:  "Illegal",
			expected: Token{
				Type:  Illegal,
				Value: "Illegal",
			},
		},
		{
			name: "Eof",
			t:    Eof,
			lit:  "Eof",
			expected: Token{
				Type:  Eof,
				Value: "Eof",
			},
		},
		{
			name: "Ident",
			t:    Ident,
			lit:  "Ident",
			expected: Token{
				Type:  Ident,
				Value: "Ident",
			},
		},
		{
			name: "Assign",
			t:    Assign,
			lit:  "Assign",
			expected: Token{
				Type:  Assign,
				Value: "Assign",
			},
		},
		{
			name: "Plus",
			t:    Plus,
			lit:  "Plus",
			expected: Token{
				Type:  Plus,
				Value: "Plus",
			},
		},
		{
			name: "Minus",
			t:    Minus,
			lit:  "Minus",
			expected: Token{
				Type:  Minus,
				Value: "Minus",
			},
		},
		{
			name: "Divide",
			t:    Divide,
			lit:  "Divide",
			expected: Token{
				Type:  Divide,
				Value: "Divide",
			},
		},
		{
			name: "Multiply",
			t:    Multiply,
			lit:  "Multiply",
			expected: Token{
				Type:  Multiply,
				Value: "Multiply",
			},
		},
		{
			name: "Modulus",
			t:    Modulus,
			lit:  "Modulus",
			expected: Token{
				Type:  Modulus,
				Value: "Modulus",
			},
		},
		{
			name: "Comma",
			t:    Comma,
			lit:  "Comma",
			expected: Token{
				Type:  Comma,
				Value: "Comma",
			},
		},
		{
			name: "FullStop",
			t:    FullStop,
			lit:  "FullStop",
			expected: Token{
				Type:  FullStop,
				Value: "FullStop",
			},
		},
		{
			name: "Semicolon",
			t:    Semicolon,
			lit:  "Semicolon",
			expected: Token{
				Type:  Semicolon,
				Value: "Semicolon",
			},
		},
		{
			name: "Colon",
			t:    Colon,
			lit:  "Colon",
			expected: Token{
				Type:  Colon,
				Value: "Colon",
			},
		},
		{
			name: "LParen",
			t:    LParen,
			lit:  "LParen",
			expected: Token{
				Type:  LParen,
				Value: "LParen",
			},
		},
		{
			name: "RParen",
			t:    RParen,
			lit:  "RParen",
			expected: Token{
				Type:  RParen,
				Value: "RParen",
			},
		},
		{
			name: "LBrace",
			t:    LBrace,
			lit:  "LBrace",
			expected: Token{
				Type:  LBrace,
				Value: "LBrace",
			},
		},
		{
			name: "RBrace",
			t:    RBrace,
			lit:  "RBrace",
			expected: Token{
				Type:  RBrace,
				Value: "RBrace",
			},
		},
		{
			name: "LBracket",
			t:    LBracket,
			lit:  "LBracket",
			expected: Token{
				Type:  LBracket,
				Value: "LBracket",
			},
		},
		{
			name: "RBracket",
			t:    RBracket,
			lit:  "RBracket",
			expected: Token{
				Type:  RBracket,
				Value: "RBracket",
			},
		},
		{
			name: "Percent",
			t:    Percent,
			lit:  "Percent",
			expected: Token{
				Type:  Percent,
				Value: "Percent",
			},
		},
		{
			name: "DoubleQuote",
			t:    DoubleQuote,
			lit:  "DoubleQuote",
			expected: Token{
				Type:  DoubleQuote,
				Value: "DoubleQuote",
			},
		},
		{
			name: "SingleQuote",
			t:    SingleQuote,
			lit:  "SingleQuote",
			expected: Token{
				Type:  SingleQuote,
				Value: "SingleQuote",
			},
		},
		{
			name: "GreaterThan",
			t:    GreaterThan,
			lit:  "GreaterThan",
			expected: Token{
				Type:  GreaterThan,
				Value: "GreaterThan",
			},
		},
		{
			name: "LessThan",
			t:    LessThan,
			lit:  "LessThan",
			expected: Token{
				Type:  LessThan,
				Value: "LessThan",
			},
		},
		{
			name: "GreaterThanOrEqual",
			t:    GreaterThanOrEqual,
			lit:  "GreaterThanOrEqual",
			expected: Token{
				Type:  GreaterThanOrEqual,
				Value: "GreaterThanOrEqual",
			},
		},
		{
			name: "LessThanOrEqual",
			t:    LessThanOrEqual,
			lit:  "LessThanOrEqual",
			expected: Token{
				Type:  LessThanOrEqual,
				Value: "LessThanOrEqual",
			},
		},
		{
			name: "Equal",
			t:    Equal,
			lit:  "Equal",
			expected: Token{
				Type:  Equal,
				Value: "Equal",
			},
		},
		{
			name: "NotEqual",
			t:    NotEqual,
			lit:  "NotEqual",
			expected: Token{
				Type:  NotEqual,
				Value: "NotEqual",
			},
		},
		{
			name: "And",
			t:    And,
			lit:  "And",
			expected: Token{
				Type:  And,
				Value: "And",
			},
		},
		{
			name: "Or",
			t:    Or,
			lit:  "Or",
			expected: Token{
				Type:  Or,
				Value: "Or",
			},
		},
		{
			name: "Not",
			t:    Not,
			lit:  "Not",
			expected: Token{
				Type:  Not,
				Value: "Not",
			},
		},
		{
			name: "Number",
			t:    Number,
			lit:  "Number",
			expected: Token{
				Type:  Number,
				Value: "Number",
			},
		},
		{
			name: "String",
			t:    String,
			lit:  "String",
			expected: Token{
				Type:  String,
				Value: "String",
			},
		},
		{
			name: "If",
			t:    If,
			lit:  "If",
			expected: Token{
				Type:  If,
				Value: "If",
			},
		},
		{
			name: "Else",
			t:    Else,
			lit:  "Else",
			expected: Token{
				Type:  Else,
				Value: "Else",
			},
		},
		{
			name: "While",
			t:    While,
			lit:  "While",
			expected: Token{
				Type:  While,
				Value: "While",
			},
		},
		{
			name: "Do",
			t:    Do,
			lit:  "Do",
			expected: Token{
				Type:  Do,
				Value: "Do",
			},
		},
		{
			name: "For",
			t:    For,
			lit:  "For",
			expected: Token{
				Type:  For,
				Value: "For",
			},
		},
		{
			name: "Function",
			t:    Function,
			lit:  "Function",
			expected: Token{
				Type:  Function,
				Value: "Function",
			},
		},
		{
			name: "Define",
			t:    Define,
			lit:  "Define",
			expected: Token{
				Type:  Define,
				Value: "Define",
			},
		},
		{
			name: "Const",
			t:    Const,
			lit:  "Const",
			expected: Token{
				Type:  Const,
				Value: "Const",
			},
		},
		{
			name: "Class",
			t:    Class,
			lit:  "Class",
			expected: Token{
				Type:  Class,
				Value: "Class",
			},
		},
		{
			name: "Include",
			t:    Include,
			lit:  "Include",
			expected: Token{
				Type:  Include,
				Value: "Include",
			},
		},
		{
			name: "Interface",
			t:    Interface,
			lit:  "Interface",
			expected: Token{
				Type:  Interface,
				Value: "Interface",
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
		lit      string
		expected string
	}{
		{
			name:     "Illegal",
			t:        Illegal,
			lit:      "Illegal",
			expected: "[ type: Illegal, value: Illegal ]",
		},
		{
			name:     "Eof",
			t:        Eof,
			lit:      "Eof",
			expected: "[ type: Eof, value: Eof ]",
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
		t        Token
		expected string
	}{
		{
			name: "Illegal",
			t: Token{
				Type:  Illegal,
				Value: "Illegal",
			},
			expected: "Illegal",
		},
		{
			name: "Eof",
			t: Token{
				Type:  Eof,
				Value: "Eof",
			},
			expected: "Eof",
		},
		{
			name: "Ident",
			t: Token{
				Type:  Ident,
				Value: "Ident",
			},
			expected: "Ident",
		},
		{
			name: "Assign",
			t: Token{
				Type:  Assign,
				Value: "Assign",
			},
			expected: "Assign",
		},
		{
			name: "Plus",
			t: Token{
				Type:  Plus,
				Value: "Plus",
			},
			expected: "Plus",
		},
		{
			name: "Minus",
			t: Token{
				Type:  Minus,
				Value: "Minus",
			},
			expected: "Minus",
		},
		{
			name: "Divide",
			t: Token{
				Type:  Divide,
				Value: "Divide",
			},
			expected: "Divide",
		},
		{
			name: "Multiply",
			t: Token{
				Type:  Multiply,
				Value: "Multiply",
			},
			expected: "Multiply",
		},
		{
			name: "Comma",
			t: Token{
				Type:  Comma,
				Value: "Comma",
			},
			expected: "Comma",
		},
		{
			name: "FullStop",
			t: Token{
				Type:  FullStop,
				Value: "FullStop",
			},
			expected: "FullStop",
		},
		{
			name: "Semicolon",
			t: Token{
				Type:  Semicolon,
				Value: "Semicolon",
			},
			expected: "Semicolon",
		},
		{
			name: "Colon",
			t: Token{
				Type:  Colon,
				Value: "Colon",
			},
			expected: "Colon",
		},
		{
			name: "LParen",
			t: Token{
				Type:  LParen,
				Value: "LParen",
			},
			expected: "LParen",
		},
		{
			name: "RParen",
			t: Token{
				Type:  RParen,
				Value: "RParen",
			},
			expected: "RParen",
		},
		{
			name: "LBrace",
			t: Token{
				Type:  LBrace,
				Value: "LBrace",
			},
			expected: "LBrace",
		},
		{
			name: "RBrace",
			t: Token{
				Type:  RBrace,
				Value: "RBrace",
			},
			expected: "RBrace",
		},
		{
			name: "LBracket",
			t: Token{
				Type:  LBracket,
				Value: "LBracket",
			},
			expected: "LBracket",
		},
		{
			name: "RBracket",
			t: Token{
				Type:  RBracket,
				Value: "RBracket",
			},
			expected: "RBracket",
		},
		{
			name: "Percent",
			t: Token{
				Type:  Percent,
				Value: "Percent",
			},
			expected: "Percent",
		},
		{
			name: "DoubleQuote",
			t: Token{
				Type:  DoubleQuote,
				Value: "DoubleQuote",
			},
			expected: "DoubleQuote",
		},
		{
			name: "SingleQuote",
			t: Token{
				Type:  SingleQuote,
				Value: "SingleQuote",
			},
			expected: "SingleQuote",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GetStringFromTokenType(tt.t)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected: %v, Actual: %v", tt.expected, actual)
			}
		})
	}

}
