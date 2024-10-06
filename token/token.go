/*
Package token defines the token type and the token struct.
*/
package token

import "fmt"

// Type is a string that represents the type of the token
type Type string

// Token is the struct that holds the token information
type Token struct {
	Type    Type        // The type of the toke. See the constants below
	Lexeme  string      // The actual string of the token
	Literal interface{} // The literal value of the token
	Line    int         // Line number where the token was found
	Column  int         // Column number where the token was found
}

//nolint:revive,stylecheck // Constants are in uppercase
const (
	// Single-character tokens
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"
	COMMA       = ","
	DOT         = "."
	MINUS       = "-"
	PLUS        = "+"
	SEMICOLON   = ";"
	SLASH       = "/"
	STAR        = "*"
	QUESTION    = "?"
	COLON       = ":"

	// One or two character tokens
	BANG          = "!"
	BANG_EQUAL    = "!="
	EQUAL         = "="
	EQUAL_EQUAL   = "=="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="

	// Literals
	IDENTIFIER = "IDENTIFIER"
	STRING     = "STRING"
	NUMBER     = "NUMBER"

	// Keywords
	AND    = "AND"
	CLASS  = "CLASS"
	ELSE   = "ELSE"
	FALSE  = "FALSE"
	FUN    = "FUN"
	FOR    = "FOR"
	IF     = "IF"
	NULL   = "NULL"
	OR     = "OR"
	PRINT  = "PRINT"
	RETURN = "RETURN"
	SUPER  = "SUPER"
	THIS   = "THIS"
	TRUE   = "TRUE"
	VAR    = "VAR"
	WHILE  = "WHILE"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

// Keywords is a map of all the reserved keywords in the language
var Keywords = map[string]Type{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"null":   NULL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

// New creates a new token
func New(t Type, l string, lit interface{}, line int) *Token {
	return &Token{Type: t, Lexeme: l, Literal: lit, Line: line}
}

// String returns the string representation of the token
func (t *Token) String() string {
	return fmt.Sprintf("%v %v %v", t.Type, t.Lexeme, t.Literal)
}
