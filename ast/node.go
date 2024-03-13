package ast

type Node interface {
	TokenValue() string
	String() string
}
