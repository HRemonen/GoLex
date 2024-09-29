/*
Go Lex programming language is a simple programming language that is
following the book "Crafting Interpreters" by Bob Nystrom. This is a
learning project for me to understand how interpreters work and how to
write one.
*/
package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Go Lex programming language!\n",
		user.Username)
}
