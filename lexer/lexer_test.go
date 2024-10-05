package lexer

import (
	"golox/token"
	"reflect"
	"testing"
)

func TestScanTokens_Characters(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedTokens []token.Token
	}{
		{
			name:  "Single character tokens",
			input: "() {} . , - + ; / * ? :",
			expectedTokens: []token.Token{
				{Type: token.LEFT_PAREN, Lexeme: "(", Literal: nil, Line: 1, Column: 1},
				{Type: token.RIGHT_PAREN, Lexeme: ")", Literal: nil, Line: 1, Column: 2},
				{Type: token.LEFT_BRACE, Lexeme: "{", Literal: nil, Line: 1, Column: 4},
				{Type: token.RIGHT_BRACE, Lexeme: "}", Literal: nil, Line: 1, Column: 5},
				{Type: token.DOT, Lexeme: ".", Literal: nil, Line: 1, Column: 7},
				{Type: token.COMMA, Lexeme: ",", Literal: nil, Line: 1, Column: 9},
				{Type: token.MINUS, Lexeme: "-", Literal: nil, Line: 1, Column: 11},
				{Type: token.PLUS, Lexeme: "+", Literal: nil, Line: 1, Column: 13},
				{Type: token.SEMICOLON, Lexeme: ";", Literal: nil, Line: 1, Column: 15},
				{Type: token.SLASH, Lexeme: "/", Literal: nil, Line: 1, Column: 17},
				{Type: token.STAR, Lexeme: "*", Literal: nil, Line: 1, Column: 19},
				{Type: token.QUESTION, Lexeme: "?", Literal: nil, Line: 1, Column: 21},
				{Type: token.COLON, Lexeme: ":", Literal: nil, Line: 1, Column: 23},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 24},
			},
		},
		{
			name:  "Unrecognized characters",
			input: "@#^",
			expectedTokens: []token.Token{
				{Type: token.ILLEGAL, Lexeme: "@", Literal: nil, Line: 1, Column: 1},
				{Type: token.ILLEGAL, Lexeme: "#", Literal: nil, Line: 1, Column: 2},
				{Type: token.ILLEGAL, Lexeme: "^", Literal: nil, Line: 1, Column: 3},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 4},
			},
		},
		{
			name:  "Comment withouth newline",
			input: "// This is a comment",
			expectedTokens: []token.Token{
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 21},
			},
		},
		{
			name:  "Comment with newline",
			input: "// This is also a comment\n",
			expectedTokens: []token.Token{
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 2, Column: 1},
			},
		},
		{
			name:  "Block comment",
			input: "/* This is a block comment */",
			expectedTokens: []token.Token{
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 30},
			},
		},
		{
			name: "Multiline block comment",
			input: `/*
				This is a multiline
				block comment
				*/
			`,
			expectedTokens: []token.Token{
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 5, Column: 4},
			},
		},
		{
			name:  "Unterminated block comment",
			input: "/* This is an unterminated block comment",
			expectedTokens: []token.Token{
				{Type: token.ILLEGAL, Lexeme: "/* This is an unterminated block comment", Literal: nil, Line: 1, Column: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 41},
			},
		},
		{
			name:  "One or two character operators",
			input: "! != = == < <= > >=",
			expectedTokens: []token.Token{
				{Type: token.BANG, Lexeme: "!", Literal: nil, Line: 1, Column: 1},
				{Type: token.BANG_EQUAL, Lexeme: "!=", Literal: nil, Line: 1, Column: 3},
				{Type: token.EQUAL, Lexeme: "=", Literal: nil, Line: 1, Column: 6},
				{Type: token.EQUAL_EQUAL, Lexeme: "==", Literal: nil, Line: 1, Column: 8},
				{Type: token.LESS, Lexeme: "<", Literal: nil, Line: 1, Column: 11},
				{Type: token.LESS_EQUAL, Lexeme: "<=", Literal: nil, Line: 1, Column: 13},
				{Type: token.GREATER, Lexeme: ">", Literal: nil, Line: 1, Column: 16},
				{Type: token.GREATER_EQUAL, Lexeme: ">=", Literal: nil, Line: 1, Column: 18},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 20},
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
				{Type: token.STRING, Lexeme: `"hello"`, Literal: "hello", Line: 1, Column: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 8},
			},
		},
		{
			name:  "STRING: Empty string",
			input: `""`,
			expectedTokens: []token.Token{
				{Type: token.STRING, Lexeme: `""`, Literal: "", Line: 1, Column: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 3},
			},
		},
		{
			name:  "STRING: Unterminated string",
			input: `"hello`,
			expectedTokens: []token.Token{
				{Type: token.ILLEGAL, Lexeme: `"hello`, Literal: nil, Line: 1, Column: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 7},
			},
		},
		{
			name:  "NUMBER: Integer",
			input: "123",
			expectedTokens: []token.Token{
				{Type: token.NUMBER, Lexeme: "123", Literal: 123.0, Line: 1, Column: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 4},
			},
		},
		{
			name:  "NUMBER: Float",
			input: "123.45",
			expectedTokens: []token.Token{
				{Type: token.NUMBER, Lexeme: "123.45", Literal: 123.45, Line: 1, Column: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 7},
			},
		},
		{
			name:  "IDENTIFIER: Single character",
			input: "a",
			expectedTokens: []token.Token{
				{Type: token.IDENTIFIER, Lexeme: "a", Literal: nil, Line: 1, Column: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 2},
			},
		},
		{
			name:  "IDENTIFIER: Multiple characters",
			input: "abc",
			expectedTokens: []token.Token{
				{Type: token.IDENTIFIER, Lexeme: "abc", Literal: nil, Line: 1, Column: 1},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 4},
			},
		},
		{
			name:  "IDENTIFIER: Keywords",
			input: "and class else false for fun if null or print return super this true var while",
			expectedTokens: []token.Token{
				{Type: token.AND, Lexeme: "and", Literal: nil, Line: 1, Column: 1},
				{Type: token.CLASS, Lexeme: "class", Literal: nil, Line: 1, Column: 5},
				{Type: token.ELSE, Lexeme: "else", Literal: nil, Line: 1, Column: 11},
				{Type: token.FALSE, Lexeme: "false", Literal: nil, Line: 1, Column: 16},
				{Type: token.FOR, Lexeme: "for", Literal: nil, Line: 1, Column: 22},
				{Type: token.FUN, Lexeme: "fun", Literal: nil, Line: 1, Column: 26},
				{Type: token.IF, Lexeme: "if", Literal: nil, Line: 1, Column: 30},
				{Type: token.NULL, Lexeme: "null", Literal: nil, Line: 1, Column: 33},
				{Type: token.OR, Lexeme: "or", Literal: nil, Line: 1, Column: 38},
				{Type: token.PRINT, Lexeme: "print", Literal: nil, Line: 1, Column: 41},
				{Type: token.RETURN, Lexeme: "return", Literal: nil, Line: 1, Column: 47},
				{Type: token.SUPER, Lexeme: "super", Literal: nil, Line: 1, Column: 54},
				{Type: token.THIS, Lexeme: "this", Literal: nil, Line: 1, Column: 60},
				{Type: token.TRUE, Lexeme: "true", Literal: nil, Line: 1, Column: 65},
				{Type: token.VAR, Lexeme: "var", Literal: nil, Line: 1, Column: 70},
				{Type: token.WHILE, Lexeme: "while", Literal: nil, Line: 1, Column: 74},
				{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1, Column: 79},
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
