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
		ln       int
		col      int
		expected Token
	}{
		{
			name: "Illegal",
			t:    Illegal,
			lit:  "Illegal",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Illegal,
				Value:  "Illegal",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Eof",
			t:    Eof,
			lit:  "Eof",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Eof,
				Value:  "Eof",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Ident",
			t:    Ident,
			lit:  "Ident",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Ident,
				Value:  "Ident",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Assign",
			t:    Assign,
			lit:  "Assign",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Assign,
				Value:  "Assign",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Plus",
			t:    Plus,
			lit:  "Plus",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Plus,
				Value:  "Plus",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Minus",
			t:    Minus,
			lit:  "Minus",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Minus,
				Value:  "Minus",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Divide",
			t:    Divide,
			lit:  "Divide",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Divide,
				Value:  "Divide",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Multiply",
			t:    Multiply,
			lit:  "Multiply",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Multiply,
				Value:  "Multiply",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Modulus",
			t:    Modulus,
			lit:  "Modulus",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Modulus,
				Value:  "Modulus",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Comma",
			t:    Comma,
			lit:  "Comma",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Comma,
				Value:  "Comma",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "FullStop",
			t:    FullStop,
			lit:  "FullStop",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   FullStop,
				Value:  "FullStop",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Semicolon",
			t:    Semicolon,
			lit:  "Semicolon",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Semicolon,
				Value:  "Semicolon",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Colon",
			t:    Colon,
			lit:  "Colon",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Colon,
				Value:  "Colon",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "LParen",
			t:    LParen,
			lit:  "LParen",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   LParen,
				Value:  "LParen",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "RParen",
			t:    RParen,
			lit:  "RParen",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   RParen,
				Value:  "RParen",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "LBrace",
			t:    LBrace,
			lit:  "LBrace",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   LBrace,
				Value:  "LBrace",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "RBrace",
			t:    RBrace,
			lit:  "RBrace",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   RBrace,
				Value:  "RBrace",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "LBracket",
			t:    LBracket,
			lit:  "LBracket",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   LBracket,
				Value:  "LBracket",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "RBracket",
			t:    RBracket,
			lit:  "RBracket",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   RBracket,
				Value:  "RBracket",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Percent",
			t:    Percent,
			lit:  "Percent",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Percent,
				Value:  "Percent",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "DoubleQuote",
			t:    DoubleQuote,
			lit:  "DoubleQuote",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   DoubleQuote,
				Value:  "DoubleQuote",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "SingleQuote",
			t:    SingleQuote,
			lit:  "SingleQuote",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   SingleQuote,
				Value:  "SingleQuote",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "GreaterThan",
			t:    GreaterThan,
			lit:  "GreaterThan",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   GreaterThan,
				Value:  "GreaterThan",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "LessThan",
			t:    LessThan,
			lit:  "LessThan",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   LessThan,
				Value:  "LessThan",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "GreaterThanOrEqual",
			t:    GreaterThanOrEqual,
			lit:  "GreaterThanOrEqual",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   GreaterThanOrEqual,
				Value:  "GreaterThanOrEqual",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "LessThanOrEqual",
			t:    LessThanOrEqual,
			lit:  "LessThanOrEqual",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   LessThanOrEqual,
				Value:  "LessThanOrEqual",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Equal",
			t:    Equal,
			lit:  "Equal",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Equal,
				Value:  "Equal",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "NotEqual",
			t:    NotEqual,
			lit:  "NotEqual",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   NotEqual,
				Value:  "NotEqual",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "And",
			t:    And,
			lit:  "And",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   And,
				Value:  "And",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Or",
			t:    Or,
			lit:  "Or",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Or,
				Value:  "Or",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Not",
			t:    Not,
			lit:  "Not",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Not,
				Value:  "Not",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Number",
			t:    Number,
			lit:  "Number",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Number,
				Value:  "Number",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "String",
			t:    String,
			lit:  "String",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   String,
				Value:  "String",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "If",
			t:    If,
			lit:  "If",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   If,
				Value:  "If",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Else",
			t:    Else,
			lit:  "Else",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Else,
				Value:  "Else",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "While",
			t:    While,
			lit:  "While",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   While,
				Value:  "While",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Do",
			t:    Do,
			lit:  "Do",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Do,
				Value:  "Do",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "For",
			t:    For,
			lit:  "For",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   For,
				Value:  "For",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Function",
			t:    Function,
			lit:  "Function",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Function,
				Value:  "Function",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Define",
			t:    Define,
			lit:  "Define",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Define,
				Value:  "Define",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Const",
			t:    Const,
			lit:  "Const",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Const,
				Value:  "Const",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Let",
			t:    Let,
			lit:  "Let",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Let,
				Value:  "Let",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Class",
			t:    Class,
			lit:  "Class",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Class,
				Value:  "Class",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Include",
			t:    Include,
			lit:  "Include",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Include,
				Value:  "Include",
				Line:   1,
				Column: 2,
			},
		},
		{
			name: "Interface",
			t:    Interface,
			lit:  "Interface",
			ln:   1,
			col:  2,
			expected: Token{
				Type:   Interface,
				Value:  "Interface",
				Line:   1,
				Column: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := New(tt.t, tt.lit, tt.ln, tt.col)
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
		ln       int
		col      int
		expected string
	}{
		{
			name:     "Illegal",
			t:        Illegal,
			lit:      "Illegal",
			ln:       1,
			col:      2,
			expected: "[ type: Illegal, value: Illegal, position: 1:2 ]",
		},
		{
			name:     "Eof",
			t:        Eof,
			lit:      "Eof",
			ln:       1,
			col:      2,
			expected: "[ type: Eof, value: Eof, position: 1:2 ]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := New(tt.t, tt.lit, tt.ln, tt.col).String()
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
			actual := GetStringFromTokenType(tt.t.Type)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected: %v, Actual: %v", tt.expected, actual)
			}
		})
	}

}
