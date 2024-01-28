package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, this is a test"
	currentPosition := 0

	runeAtIndex, size := utf8.DecodeRuneInString(str[currentPosition:])

	if runeAtIndex != utf8.RuneError || size > 0 {
		fmt.Printf("Rune at index %d: %c\n", currentPosition, runeAtIndex)
	} else {
		fmt.Printf("Invalid utf-8 sequence at index %d\n", currentPosition)
	}
}
