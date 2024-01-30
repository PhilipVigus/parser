package main

import (
	"fmt"
	"parser/lexer"
	"parser/lexer/token"
)

func main() {
	l, err := lexer.New("bob-+wibble")

	if err == nil {
		for {
			t := l.NextToken()

			fmt.Println(t)

			if t.Type == token.EOF || t.Type == token.ILLEGAL {
				break
			}
		}
	}
}
