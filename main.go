/*
Go Lox programming language is a simple programming language that is
following the book "Crafting Interpreters" by Bob Nystrom. This is a
learning project for me to understand how interpreters work and how to
write one.
*/
package main

import (
	"fmt"
	"golox/repl"
	"os"
)

func main() {
	fmt.Println("Welcome to GoLox!\n Feel free to type in commands")

	repl.Start(os.Stdin, os.Stdout)
}
