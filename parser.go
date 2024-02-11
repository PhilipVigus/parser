package main

import (
	"fmt"
	"parser/lexer"
	"strings"
)

func main() {
	inputString := "it;"
	lexerFromString := lexer.New(strings.NewReader(inputString))

	tokens := lexerFromString.Tokenize()

	fmt.Println(tokens)
}
