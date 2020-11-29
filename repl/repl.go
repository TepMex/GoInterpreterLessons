package repl

import (
	"bufio"
	"fmt"
	"interpreterlesson/lexer"
	"interpreterlesson/token"
	"io"
)

// PROMPT - REPL invitation
const PROMPT = ">>> "

// Start - starts a REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
