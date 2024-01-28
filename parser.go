package main

import (
	"fmt"
	"parser/lexer"
	"parser/lexer/token"
)

func main() {
	l, err := lexer.New("+-")

	if err == nil {
		for {
			t := l.NextToken()

			fmt.Printf("%d : %d", t.Type, t.Literal)

			if t.Type == token.EOF || t.Type == token.ILLEGAL {
				break
			}
		}
	}
}
