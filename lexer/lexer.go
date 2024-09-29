/*
Package lexer implements the lexer for the Lox programming language. The lexer
is responsible for scanning the source code and converting it into a list of
tokens that the parser can consume.
*/
package lexer

import (
	"fmt"
	"golox/token"
	"strconv"
)

// Lexer is the struct that holds the state of the lexer
type Lexer struct {
	source  string
	Tokens  []token.Token
	start   int // Start of the current lexeme
	current int // Current character being looked at
	line    int // Current line number
}

// New creates a new lexer
func New(source string) *Lexer {
	return &Lexer{
		source:  source,
		Tokens:  []token.Token{},
		start:   0,
		current: 0,
		line:    1,
	}
}

// isAtEnd checks if the lexer has reached the end of the source code
func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

func (l *Lexer) isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (l *Lexer) isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (l *Lexer) isAlphaNumeric(c rune) bool {
	return l.isAlpha(c) || l.isDigit(c)
}

// match checks if the current character matches the expected character
func (l *Lexer) match(expected rune) bool {
	if l.isAtEnd() {
		return false
	}

	if rune(l.source[l.current]) != expected {
		return false
	}

	l.current++
	return true
}

// peek returns the next character in the source code
func (l *Lexer) peek() rune {
	if l.isAtEnd() {
		return '\x00'
	}

	return rune(l.source[l.current])
}

// peekNext returns the character after the next character in the source code
func (l *Lexer) peekNext() rune {
	if l.current+1 >= len(l.source) {
		return '\x00'
	}

	return rune(l.source[l.current+1])
}

func (l *Lexer) string() (string, error) {
	for l.peek() != '"' && !l.isAtEnd() {
		if l.peek() == '\n' {
			l.line++
		}
		l.advance()
	}

	if l.isAtEnd() {
		return "", fmt.Errorf("unterminated string")
	}

	// The closing "
	l.advance()

	// Trim the surrounding quotes
	value := l.source[l.start+1 : l.current-1]
	return value, nil
}

func (l *Lexer) number() (float64, error) {
	for l.isDigit(l.peek()) {
		l.advance()
	}

	// Look for a fractional part
	if l.peek() == '.' && l.isDigit(l.peekNext()) {
		// Consume the "."
		l.advance()

		for l.isDigit(l.peek()) {
			l.advance()
		}
	}

	value := l.source[l.start:l.current]
	return strconv.ParseFloat(value, 64)
}

func (l *Lexer) identifier() token.Type {
	for l.isAlphaNumeric(l.peek()) {
		l.advance()
	}

	text := l.source[l.start:l.current]

	tokenType, ok := token.Keywords[text]
	if !ok {
		tokenType = token.IDENTIFIER
	}

	return tokenType
}

//nolint:funlen,gocyclo // This function is long and complex because it has to handle all the different token types
func (l *Lexer) scanToken() {
	c := l.advance()

	switch c {
	case '(':
		l.addToken(token.LEFT_PAREN, nil)
	case ')':
		l.addToken(token.RIGHT_PAREN, nil)
	case '{':
		l.addToken(token.LEFT_BRACE, nil)
	case '}':
		l.addToken(token.RIGHT_BRACE, nil)
	case ',':
		l.addToken(token.COMMA, nil)
	case '.':
		l.addToken(token.DOT, nil)
	case '-':
		l.addToken(token.MINUS, nil)
	case '+':
		l.addToken(token.PLUS, nil)
	case ';':
		l.addToken(token.SEMICOLON, nil)
	case '/':
		if l.match('/') {
			// A comment goes until the end of the line
			for l.peek() != '\n' && !l.isAtEnd() {
				l.advance()
			}
			break
		}
		l.addToken(token.SLASH, nil)
	case '*':
		l.addToken(token.STAR, nil)
	case '!':
		if l.match('=') {
			l.addToken(token.BANG_EQUAL, nil)
			break
		}
		l.addToken(token.BANG, nil)
	case '=':
		if l.match('=') {
			l.addToken(token.EQUAL_EQUAL, nil)
			break
		}
		l.addToken(token.EQUAL, nil)
	case '<':
		if l.match('=') {
			l.addToken(token.LESS_EQUAL, nil)
			break
		}
		l.addToken(token.LESS, nil)
	case '>':
		if l.match('=') {
			l.addToken(token.GREATER_EQUAL, nil)
			break
		}
		l.addToken(token.GREATER, nil)
	case ' ', '\r', '\t': // Ignore whitespace
	case '\n': // Ignore newline
		l.line++
	case '"': // String literals
		value, err := l.string()
		if err != nil {
			l.addToken(token.ILLEGAL, nil)
			break
		}
		l.addToken(token.STRING, value)
	default:
		if l.isDigit(c) {
			value, err := l.number()
			if err != nil {
				l.addToken(token.ILLEGAL, nil)
				break
			}

			l.addToken(token.NUMBER, value)
			break
		} else if l.isAlpha(c) {
			value := l.identifier()

			l.addToken(value, nil)
			break
		}

		l.addToken(token.ILLEGAL, nil)
	}
}

func (l *Lexer) advance() rune {
	l.current++
	return rune(l.source[l.current-1])
}

func (l *Lexer) addToken(tokenType token.Type, literal interface{}) {
	text := l.source[l.start:l.current]
	l.Tokens = append(l.Tokens, token.Token{
		Type:    tokenType,
		Lexeme:  text,
		Literal: literal,
		Line:    l.line,
	})
}

// ScanTokens scans the source code and converts it into a list of tokens
func (l *Lexer) ScanTokens() {
	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}

	// Add EOF token to the end of the tokens list
	l.Tokens = append(l.Tokens, token.Token{
		Type:    token.EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    l.line,
	})
}
