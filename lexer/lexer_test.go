package lexer

import (
	"reflect"
	"testing"

	"golox/token"
)

func TestScanTokens_SingleCharacters(t *testing.T) {
	tests := []struct {
		input          string
		expectedTokens []token.Token
	}{
		{
			input: "(){}.,-+;/*",
			expectedTokens: []token.Token{
				{Type: token.LEFT_PAREN, Lexeme: "(", Literal: nil, Line: 1},
				{Type: token.RIGHT_PAREN, Lexeme: ")", Literal: nil, Line: 1},
				{Type: token.LEFT_BRACE, Lexeme: "{", Literal: nil, Line: 1},
				{Type: token.RIGHT_BRACE, Lexeme: "}", Literal: nil, Line: 1},
				{Type: token.DOT, Lexeme: ".", Literal: nil, Line: 1},
				{Type: token.COMMA, Lexeme: ",", Literal: nil, Line: 1},
				{Type: token.MINUS, Lexeme: "-", Literal: nil, Line: 1},
				{Type: token.PLUS, Lexeme: "+", Literal: nil, Line: 1},
				{Type: token.SEMICOLON, Lexeme: ";", Literal: nil, Line: 1},
				{Type: token.SLASH, Lexeme: "/", Literal: nil, Line: 1},
				{Type: token.STAR, Lexeme: "*", Literal: nil, Line: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
		{
			input: "@#^",
			expectedTokens: []token.Token{
				{Type: token.ILLEGAL, Lexeme: "@", Literal: nil, Line: 1},
				{Type: token.ILLEGAL, Lexeme: "#", Literal: nil, Line: 1},
				{Type: token.ILLEGAL, Lexeme: "^", Literal: nil, Line: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
		{
			input: "// This is a comment",
			expectedTokens: []token.Token{
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
	}

	for _, tt := range tests {
		l := New(tt.input)
		l.scanTokens()
		if !reflect.DeepEqual(l.tokens, tt.expectedTokens) {
			t.Errorf("For input %q, expected tokens %v, but got %v", tt.input, tt.expectedTokens, l.tokens)
		}
	}
}
