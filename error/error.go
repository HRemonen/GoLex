package error

import (
	"fmt"
	"golox/token"
)

// Error represents an error
type Error struct {
	Message string
	Token   *token.Token
}

// NewError creates a new error
func New(t *token.Token, message string) *Error {
	return &Error{
		Message: message,
		Token:   t,
	}
}

func (e *Error) Error() string {
	if e.Token.Type == token.EOF {
		return fmt.Sprintf("[Pos %d:%d] Error at end: %s", e.Token.Line, e.Token.Column, e.Message)
	}

	return fmt.Sprintf("[Pos %d:%d] Error at '%s': %s", e.Token.Line, e.Token.Column, e.Token.Lexeme, e.Message)
}
