package repl

import (
	"bufio"
	"fmt"
	"io"
	"lang/lexer"
	"lang/lexer/token"
	"strings"
)

const Prompt = ">> "

// Start starts the REPL.
// It reads input from the given reader and writes output to the given writer.
// Output is in the form of tokens, which are generated whenever the REPL
// encounters a new line of input from the user
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, Prompt)

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		tokens := tokenize(line)

		for _, tok := range tokens {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

func tokenize(line string) []token.Token {
	l := lexer.New(strings.NewReader(line))
	return l.Tokenize()
}
