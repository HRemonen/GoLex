package lexer

import (
	"fmt"
	"golox/token"
)

// Lexer is the struct that holds the state of the lexer
type Lexer struct {
	source  string
	tokens  []token.Token
	start   int // Start of the current lexeme
	current int // Current character being looked at
	line    int // Current line number
}

// New creates a new lexer
func New(source string) *Lexer {
	return &Lexer{
		source:  source,
		tokens:  []token.Token{},
		start:   0,
		current: 0,
		line:    1,
	}
}

// isAtEnd checks if the lexer has reached the end of the source code
func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
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

func (l *Lexer) string() (string, error) {
	for l.peek() != '"' && !l.isAtEnd() {
		if l.peek() == '\n' {
			l.line++
		}
		l.advance()
	}

	fmt.Println(l.peek())

	if l.isAtEnd() {
		return "", fmt.Errorf("Unterminated string")
	}

	// The closing "
	l.advance()

	// Trim the surrounding quotes
	value := l.source[l.start+1 : l.current-1]
	fmt.Println(value)
	return value, nil
}

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
	case ' ', '\r', '\t':
		// Ignore whitespace
		break
	case '\n':
		l.line++
		// Ignore newline
		break
	case '"':
		// String literals
		value, err := l.string()
		if err != nil {
			l.addToken(token.ILLEGAL, nil)
			break
		}
		l.addToken(token.STRING, value)
	default:
		// If we reach here, it means we have an unexpected character
		l.addToken(token.ILLEGAL, nil)
	}
}

func (l *Lexer) advance() rune {
	l.current++
	return rune(l.source[l.current-1])
}

func (l *Lexer) addToken(tokenType token.TokenType, literal interface{}) {
	text := l.source[l.start:l.current]
	l.tokens = append(l.tokens, token.Token{
		Type:    tokenType,
		Lexeme:  text,
		Literal: literal,
		Line:    l.line,
	})
}

func (l *Lexer) scanTokens() {
	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}

	// Add EOF token to the end of the tokens list
	l.tokens = append(l.tokens, token.Token{
		Type:    token.EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    l.line,
	})
}
