package token

import (
	"fmt"
)

type Type int

type Token struct {
	Type   Type
	Value  string
	Line   int
	Column int
}

func New(t Type, v string, l int, c int) Token {
	return Token{
		Type:   t,
		Value:  v,
		Line:   l,
		Column: c,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("[ type: %s, value: %v, position: %d:%d ]", GetStringFromTokenType(t), t.Value, t.Line, t.Column)
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
	Let
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

var keywordsToTypes = map[string]Type{
	"if":         If,
	"else":       Else,
	"while":      While,
	"do":         Do,
	"for":        For,
	"function":   Function,
	"define":     Define,
	"const":      Const,
	"let":        Let,
	"class":      Class,
	"include":    Include,
	"interface":  Interface,
	"in":         In,
	"break":      Break,
	"continue":   Continue,
	"catch":      Catch,
	"try":        Try,
	"switch":     Switch,
	"case":       Case,
	"default":    Default,
	"enum":       Enum,
	"export":     Export,
	"new":        NewToken,
	"throw":      Throw,
	"extends":    Extends,
	"implements": Implements,
	"private":    Private,
	"protected":  Protected,
	"public":     Public,
	"static":     Static,
	"abstract":   Abstract,
	"return":     Return,
	"finally":    Finally,
}

func GetKeywordType(kw string) Type {
	t, exists := keywordsToTypes[kw]

	if exists {
		return t
	}

	return Ident
}

var tokenTypeToString = map[Type]string{
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
	Let:                "Let",
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

func GetStringFromTokenType(t Token) string {
	name, exists := tokenTypeToString[t.Type]

	if exists {
		return name
	}

	return "UNKNOWN"
}
