package lexer

import (
	"golox/token"
	"reflect"
	"testing"
)

//nolint:funlen // This is a test function
func TestScanTokens_Characters(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedTokens []token.Token
	}{
		{
			name:  "Single character tokens",
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
			name:  "Unrecognized characters",
			input: "@#^",
			expectedTokens: []token.Token{
				{Type: token.ILLEGAL, Lexeme: "@", Literal: nil, Line: 1},
				{Type: token.ILLEGAL, Lexeme: "#", Literal: nil, Line: 1},
				{Type: token.ILLEGAL, Lexeme: "^", Literal: nil, Line: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
		{
			name:  "Comment withouth newline",
			input: "// This is a comment",
			expectedTokens: []token.Token{
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
		{
			name:  "Comment with newline",
			input: "// This is also a comment\n",
			expectedTokens: []token.Token{
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 2},
			},
		},
		{
			name:  "One or two character operators",
			input: "! != = == < <= > >=",
			expectedTokens: []token.Token{
				{Type: token.BANG, Lexeme: "!", Literal: nil, Line: 1},
				{Type: token.BANG_EQUAL, Lexeme: "!=", Literal: nil, Line: 1},
				{Type: token.EQUAL, Lexeme: "=", Literal: nil, Line: 1},
				{Type: token.EQUAL_EQUAL, Lexeme: "==", Literal: nil, Line: 1},
				{Type: token.LESS, Lexeme: "<", Literal: nil, Line: 1},
				{Type: token.LESS_EQUAL, Lexeme: "<=", Literal: nil, Line: 1},
				{Type: token.GREATER, Lexeme: ">", Literal: nil, Line: 1},
				{Type: token.GREATER_EQUAL, Lexeme: ">=", Literal: nil, Line: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)

			l.ScanTokens()

			if !reflect.DeepEqual(l.Tokens, tt.expectedTokens) {
				t.Errorf("Test %s failed. Expected tokens: %v, but got: %v", tt.name, tt.expectedTokens, l.Tokens)
			}
		})
	}
}

func TestScanTokens_Literals(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedTokens []token.Token
	}{
		{
			name:  "STRING: Normal string",
			input: `"hello"`,
			expectedTokens: []token.Token{
				{Type: token.STRING, Lexeme: `"hello"`, Literal: "hello", Line: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
		{
			name:  "STRING: Empty string",
			input: `""`,
			expectedTokens: []token.Token{
				{Type: token.STRING, Lexeme: `""`, Literal: "", Line: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
		{
			name:  "STRING: Unterminated string",
			input: `"hello`,
			expectedTokens: []token.Token{
				{Type: token.ILLEGAL, Lexeme: `"hello`, Literal: nil, Line: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
		{
			name:  "NUMBER: Integer",
			input: "123",
			expectedTokens: []token.Token{
				{Type: token.NUMBER, Lexeme: "123", Literal: 123.0, Line: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
		{
			name:  "NUMBER: Float",
			input: "123.45",
			expectedTokens: []token.Token{
				{Type: token.NUMBER, Lexeme: "123.45", Literal: 123.45, Line: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)

			l.ScanTokens()

			if !reflect.DeepEqual(l.Tokens, tt.expectedTokens) {
				t.Errorf("Test %s failed. Expected tokens: %v, but got: %v", tt.name, tt.expectedTokens, l.Tokens)
			}
		})
	}
}
