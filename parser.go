package main

import (
	"fmt"
	"parser/lexer"
)

func main() {
	l := lexer.NewLexer("a test string")

	for {
		c, s := l.PeekNextChar()
		fmt.Println(string(c))

		c, s = l.GetNextChar()
		fmt.Println(string(c))

		if c == 0 {
			break
		}

		if s == 0 {
			break
		}
	}
}
