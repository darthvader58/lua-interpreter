package repl

import (
	"bufio"
	"fmt"
	"io"
	"lua-interpreter/lexer"
	"lua-interpreter/token"
)

const PROMPT = "lua> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	fmt.Fprintf(out, "Lua Interpreter (Lexer Phase)\n")
	fmt.Fprintf(out, "Type 'exit' to quit\n\n")

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			fmt.Fprintf(out, "Goodbye!\n")
			return
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
