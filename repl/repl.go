/*
Package repl provides a Read-Eval-Print-Loop for the Lox language.
*/
package repl

import (
	"bufio"
	"fmt"
	"golox/lexer"
	"golox/parser"
	"golox/printer"
	"io"
)

// PROMPT is the prompt for the REPL
const PROMPT = "> "

// Start starts the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, err := fmt.Fprint(out, PROMPT)
		if err != nil {
			fmt.Println("Error writing to output")
			return
		}

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		l.ScanTokens()

		p := parser.New(l.Tokens)
		expr := p.Parse()

		printer := printer.New()
		fmt.Println(printer.Print(expr))
	}
}
