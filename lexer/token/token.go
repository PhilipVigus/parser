package token

import "fmt"

type Type int
type Token[T any] struct {
	Type    Type
	Literal T
}

func New[T any](t Type, l T) Token[any] {
	return Token[any]{
		Type:    t,
		Literal: l,
	}
}

func (t Token[T]) String() string {
	return fmt.Sprintf("[ type: %s, literal: %v ]", GetTokenType(t), t.Literal)
}

const (
	Illegal = iota
	Eof
	Ident
	Assign
	Plus
	Minus
	Divide
	Multiply
	Modulus
	Comma
	FullStop
	Semicolon
	Colon
	LParen
	RParen
	LBrace
	RBrace
	LBracket
	RBracket
	Percent
	DoubleQuote
	SingleQuote
	GreaterThan
	LessThan
	GreaterThanOrEqual
	LessThanOrEqual
	Equal
	NotEqual
	And
	Or
	Not
	Number
	String
	If
	Else
	While
	Do
	For
	Function
	Define
	Const
	Class
	Include
	Interface
	In
	Break
	Continue
	Catch
	Try
	Switch
	Case
	Default
	Enum
	Export
	NewToken
	Throw
	Extends
	Implements
	Private
	Protected
	Public
	Static
	Abstract
	Return
	Finally
)

var tokenTypes = map[Type]string{
	Illegal:            "Illegal",
	Eof:                "Eof",
	Ident:              "Ident",
	Assign:             "Assign",
	Plus:               "Plus",
	Minus:              "Minus",
	Divide:             "Divide",
	Multiply:           "Multiply",
	Modulus:            "Modulus",
	Comma:              "Comma",
	FullStop:           "FullStop",
	Semicolon:          "Semicolon",
	Colon:              "Colon",
	LParen:             "LParen",
	RParen:             "RParen",
	LBrace:             "LBrace",
	RBrace:             "RBrace",
	LBracket:           "LBracket",
	RBracket:           "RBracket",
	Percent:            "Percent",
	DoubleQuote:        "DoubleQuote",
	SingleQuote:        "SingleQuote",
	GreaterThan:        "GreaterThan",
	LessThan:           "LessThan",
	GreaterThanOrEqual: "GreaterThanOrEqual",
	LessThanOrEqual:    "LessThanOrEqual",
	Equal:              "Equal",
	NotEqual:           "NotEqual",
	And:                "And",
	Or:                 "Or",
	Not:                "Not",
	Number:             "Number",
	String:             "String",
	If:                 "If",
	Else:               "Else",
	While:              "While",
	Do:                 "Do",
	For:                "For",
	Function:           "Function",
	Define:             "Define",
	Const:              "Const",
	Class:              "Class",
	Include:            "Include",
	Interface:          "Interface",
	In:                 "In",
	Break:              "Break",
	Continue:           "Continue",
	Catch:              "Catch",
	Try:                "Try",
	Switch:             "Switch",
	Case:               "Case",
	Default:            "Default",
	Enum:               "Enum",
	Export:             "Export",
	NewToken:           "New",
	Throw:              "Throw",
	Extends:            "Extends",
	Implements:         "Implements",
	Private:            "Private",
	Protected:          "Protected",
	Public:             "Public",
	Static:             "Static",
	Abstract:           "Abstract",
	Return:             "Return",
	Finally:            "Finally",
}

func GetTokenType[T any](t Token[T]) string {
	name, exists := tokenTypes[t.Type]

	if exists {
		return name
	}

	return "UNKNOWN"
}
