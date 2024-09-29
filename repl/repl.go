package repl

import (
	"bufio"
	"fmt"
	"io"

	"golox/lexer"
)

// PROMPT is the prompt for the REPL
const PROMPT = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		l.ScanTokens()

		for _, t := range l.Tokens {
			fmt.Println(t)
		}
	}
}
