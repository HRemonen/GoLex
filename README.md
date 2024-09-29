# Golox: Lox Interpreter in Go

**Golox**, a Lox language interpreter written in Go. This project is inspired by the book [Crafting Interpreters](https://craftinginterpreters.com/) by Robert Nystrom. The interpreter is being built incrementally, starting with lexical analysis (lexer) and progressing toward a full interpreter.

## Table of Contents
- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Linting](#linting)

## Introduction

Lox is a dynamically-typed, interpreted programming language created for learning purposes. This project, **Golox**, aims to implement a Lox interpreter using the Go programming language.

## Installation

### Prerequisites
- [Go](https://golang.org/doc/install) (version 1.23+)

### Clone the Repository

To download the source code, clone the repository:

```bash
git clone git@github.com:HRemonen/GoLox.git
cd golox
```

### Install Dependencies

Run the following command to install Go module dependencies:

```bash
go mod tidy
```

This will install any necessary packages for the project.

## Usage

Currently, the project includes a lexer that can tokenize source code written in Lox.

To use the lexer, import the package in your Go code and pass a Lox source string to the lexer.New function:

```go
package main

import (
    "fmt"
    "golox/lexer"
)

func main() {
    source := `print "Hello, world!";`
    l := lexer.New(source)
    l.ScanTokens()

    for _, token := range l.Tokens() {
        fmt.Println(token)
    }
}
```

## Testing

This project includes tests for various parts of the lexer, particularly for string literals and handling of special characters.

To run the tests, use:

```bash
go test -v ./lexer
```

## Linting

To ensure that the codebase follows Go best practices and maintain a clean, consistent style, we use `golangci-lint`, a popular linter aggregator for Go.

### Installing the Linter

First, install `golangci-lint` by following the official instructions [here](https://golangci-lint.run/usage/install/). You can also install it using `go install`:

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Running the Linter

Once installed, you can run the linter on the project using the following command:

```bash
golangci-lint run
```

This will check the entire codebase for issues and display any linting errors, warnings, or suggestions.

In some cases, the linter can automatically fix issues like formatting errors. To apply fixes automatically, run:

```bash
golangci-lint run --fix
```

The linter is also run on the CI pipeline.




