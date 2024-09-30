/*
Package lexer implements the lexer for the Lox programming language. The lexer
is responsible for scanning the source code and converting it into a list of
tokens that the parser can consume.
*/
package lexer

import (
	"golox/token"
	"strconv"
)

// Lexer holds the state of the lexer
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

// scanToken processes a single token
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
	case '*':
		l.addToken(token.STAR, nil)
	case ' ', '\r', '\t':
		// Ignore whitespace
	case '\n':
		l.line++
	case '/':
		if l.match('/') {
			l.lineComment()
		} else if l.match('*') {
			l.blockComment()
		} else {
			l.addToken(token.SLASH, nil)
		}
	case '!':
		l.addToken(l.matchToken('=', token.BANG_EQUAL, token.BANG), nil)
	case '=':
		l.addToken(l.matchToken('=', token.EQUAL_EQUAL, token.EQUAL), nil)
	case '<':
		l.addToken(l.matchToken('=', token.LESS_EQUAL, token.LESS), nil)
	case '>':
		l.addToken(l.matchToken('=', token.GREATER_EQUAL, token.GREATER), nil)
	case '"':
		l.processString()
	default:
		if l.isDigit(c) {
			l.processNumber()
		} else if l.isAlpha(c) {
			l.processIdentifier()
		} else {
			l.addIllegalToken()
		}
	}
}

// Helper for handling strings
func (l *Lexer) processString() {
	for l.peek() != '"' && !l.isAtEnd() {
		if l.peek() == '\n' {
			l.line++
		}
		l.advance()
	}

	if l.isAtEnd() {
		l.addIllegalToken()
		return
	}

	// Closing quote
	l.advance()
	value := l.source[l.start+1 : l.current-1] // Remove quotes
	l.addToken(token.STRING, value)
}

// Helper for handling numbers
func (l *Lexer) processNumber() {
	for l.isDigit(l.peek()) {
		l.advance()
	}

	// Handle fractional part
	if l.peek() == '.' && l.isDigit(l.peekNext()) {
		l.advance() // Consume '.'
		for l.isDigit(l.peek()) {
			l.advance()
		}
	}

	value := l.source[l.start:l.current]
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		l.addIllegalToken()
	} else {
		l.addToken(token.NUMBER, parsed)
	}
}

// Helper for handling identifiers and keywords
func (l *Lexer) processIdentifier() {
	for l.isAlphaNumeric(l.peek()) {
		l.advance()
	}

	text := l.source[l.start:l.current]
	tokenType, ok := token.Keywords[text]
	if !ok {
		tokenType = token.IDENTIFIER
	}
	l.addToken(tokenType, nil)
}

// Helper for handling block comments
func (l *Lexer) blockComment() {
	for !(l.peek() == '*' && l.peekNext() == '/') && !l.isAtEnd() {
		if l.peek() == '\n' {
			l.line++
		}
		l.advance()
	}

	// Consume closing '*/'
	if !l.isAtEnd() {
		l.advance() // Consume '*'
		l.advance() // Consume '/'
	} else {
		l.addIllegalToken()
	}
}

// Helper for handling single-line comments
func (l *Lexer) lineComment() {
	for l.peek() != '\n' && !l.isAtEnd() {
		l.advance()
	}
}

// Adds a token to the list
func (l *Lexer) addToken(tokenType token.Type, literal interface{}) {
	text := l.source[l.start:l.current]
	l.Tokens = append(l.Tokens, token.Token{
		Type:    tokenType,
		Lexeme:  text,
		Literal: literal,
		Line:    l.line,
	})
}

// Adds an illegal token
func (l *Lexer) addIllegalToken() {
	l.addToken(token.ILLEGAL, nil)
}

// Advances the lexer to the next character
func (l *Lexer) advance() rune {
	l.current++
	return rune(l.source[l.current-1])
}

// Matches the current character with an expected one
func (l *Lexer) match(expected rune) bool {
	if l.isAtEnd() || rune(l.source[l.current]) != expected {
		return false
	}
	l.current++
	return true
}

// Matches a token based on a condition
func (l *Lexer) matchToken(expected rune, matchType, defaultType token.Type) token.Type {
	if l.match(expected) {
		return matchType
	}
	return defaultType
}

// Peeks at the next character without advancing
func (l *Lexer) peek() rune {
	if l.isAtEnd() {
		return '\x00'
	}
	return rune(l.source[l.current])
}

// Peeks two characters ahead
func (l *Lexer) peekNext() rune {
	if l.current+1 >= len(l.source) {
		return '\x00'
	}
	return rune(l.source[l.current+1])
}

// Checks if the end of the source has been reached
func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

// Checks if the given character is a digit
func (l *Lexer) isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

// Checks if the given character is an alphabetical character or an underscore
func (l *Lexer) isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

// Checks if the character is alphanumeric or an underscore
func (l *Lexer) isAlphaNumeric(c rune) bool {
	return l.isAlpha(c) || l.isDigit(c)
}
