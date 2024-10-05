package parser

import (
	"golox/error"
	"golox/expr"
	"golox/token"

	"reflect"
	"testing"
)

func TestParser_Expressions(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []token.Token
		expected expr.Expr
	}{
		{
			name: "Simple Binary Expression (1 + 2)",
			tokens: []token.Token{
				{Type: token.NUMBER, Literal: 1},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.NUMBER, Literal: 2},
				{Type: token.EOF},
			},
			expected: &expr.Binary{
				Left:     &expr.Literal{Value: 1},
				Operator: &token.Token{Type: token.PLUS, Literal: "+"},
				Right:    &expr.Literal{Value: 2},
			},
		},
		{
			name: "Operator Precedence (1 + 2 * 3)",
			tokens: []token.Token{
				{Type: token.NUMBER, Literal: 1},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.NUMBER, Literal: 2},
				{Type: token.STAR, Literal: "*"},
				{Type: token.NUMBER, Literal: 3},
				{Type: token.EOF},
			},
			expected: &expr.Binary{
				Left:     &expr.Literal{Value: 1},
				Operator: &token.Token{Type: token.PLUS, Literal: "+"},
				Right: &expr.Binary{
					Left:     &expr.Literal{Value: 2},
					Operator: &token.Token{Type: token.STAR, Literal: "*"},
					Right:    &expr.Literal{Value: 3},
				},
			},
		},
		{
			name: "Grouping with Parentheses ((1 + 2) * 3)",
			tokens: []token.Token{
				{Type: token.LEFT_PAREN, Literal: "("},
				{Type: token.NUMBER, Literal: 1},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.NUMBER, Literal: 2},
				{Type: token.RIGHT_PAREN, Literal: ")"},
				{Type: token.STAR, Literal: "*"},
				{Type: token.NUMBER, Literal: 3},
				{Type: token.EOF},
			},
			expected: &expr.Binary{
				Left: &expr.Grouping{
					Expression: &expr.Binary{
						Left:     &expr.Literal{Value: 1},
						Operator: &token.Token{Type: token.PLUS, Literal: "+"},
						Right:    &expr.Literal{Value: 2},
					},
				},
				Operator: &token.Token{Type: token.STAR, Literal: "*"},
				Right:    &expr.Literal{Value: 3},
			},
		},
		{
			name: "Unary Expression (-5)",
			tokens: []token.Token{
				{Type: token.MINUS, Literal: "-"},
				{Type: token.NUMBER, Literal: 5},
				{Type: token.EOF},
			},
			expected: &expr.Unary{
				Operator: &token.Token{Type: token.MINUS, Literal: "-"},
				Right:    &expr.Literal{Value: 5},
			},
		},
		{
			name: "Unary Expression (!true)",
			tokens: []token.Token{
				{Type: token.BANG, Literal: "!"},
				{Type: token.TRUE, Literal: true},
				{Type: token.EOF},
			},
			expected: &expr.Unary{
				Operator: &token.Token{Type: token.BANG, Literal: "!"},
				Right:    &expr.Literal{Value: true},
			},
		},
		{
			name: "Grouping with Unary Expression (-(-5))",
			tokens: []token.Token{
				{Type: token.LEFT_PAREN, Literal: "("},
				{Type: token.MINUS, Literal: "-"},
				{Type: token.LEFT_PAREN, Literal: "("},
				{Type: token.MINUS, Literal: "-"},
				{Type: token.NUMBER, Literal: 5},
				{Type: token.RIGHT_PAREN, Literal: ")"},
				{Type: token.RIGHT_PAREN, Literal: ")"},
				{Type: token.EOF},
			},
			expected: &expr.Grouping{
				Expression: &expr.Unary{
					Operator: &token.Token{Type: token.MINUS, Literal: "-"},
					Right: &expr.Grouping{
						Expression: &expr.Unary{
							Operator: &token.Token{Type: token.MINUS, Literal: "-"},
							Right:    &expr.Literal{Value: 5},
						},
					},
				},
			},
		},
		{
			name: "Empty Expression",
			tokens: []token.Token{
				{Type: token.EOF},
			},
			expected: nil,
		},
		{
			name: "Deeply Nested Grouping (((1 + 2)))",
			tokens: []token.Token{
				{Type: token.LEFT_PAREN, Literal: "("},
				{Type: token.LEFT_PAREN, Literal: "("},
				{Type: token.LEFT_PAREN, Literal: "("},
				{Type: token.NUMBER, Literal: 1},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.NUMBER, Literal: 2},
				{Type: token.RIGHT_PAREN, Literal: ")"},
				{Type: token.RIGHT_PAREN, Literal: ")"},
				{Type: token.RIGHT_PAREN, Literal: ")"},
				{Type: token.EOF},
			},
			expected: &expr.Grouping{
				Expression: &expr.Grouping{
					Expression: &expr.Grouping{
						Expression: &expr.Binary{
							Left:     &expr.Literal{Value: 1},
							Operator: &token.Token{Type: token.PLUS, Literal: "+"},
							Right:    &expr.Literal{Value: 2},
						},
					},
				},
			},
		},
		{
			name: "Chained Binary Operators (1 + 2 + 3)",
			tokens: []token.Token{
				{Type: token.NUMBER, Literal: 1},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.NUMBER, Literal: 2},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.NUMBER, Literal: 3},
				{Type: token.EOF},
			},
			expected: &expr.Binary{
				Left: &expr.Binary{
					Left:     &expr.Literal{Value: 1},
					Operator: &token.Token{Type: token.PLUS, Literal: "+"},
					Right:    &expr.Literal{Value: 2},
				},
				Operator: &token.Token{Type: token.PLUS, Literal: "+"},
				Right:    &expr.Literal{Value: 3},
			},
		},
		{
			name: "Unary Chaining (---5)",
			tokens: []token.Token{
				{Type: token.MINUS, Literal: "-"},
				{Type: token.MINUS, Literal: "-"},
				{Type: token.MINUS, Literal: "-"},
				{Type: token.NUMBER, Literal: 5},
				{Type: token.EOF},
			},
			expected: &expr.Unary{
				Operator: &token.Token{Type: token.MINUS, Literal: "-"},
				Right: &expr.Unary{
					Operator: &token.Token{Type: token.MINUS, Literal: "-"},
					Right: &expr.Unary{
						Operator: &token.Token{Type: token.MINUS, Literal: "-"},
						Right:    &expr.Literal{Value: 5},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new parser instance with the test tokens
			p := New(tt.tokens)

			// Parse the expression
			expression := p.expression()

			// Check if the parsed expression matches the expected AST
			if !reflect.DeepEqual(expression, tt.expected) {
				t.Errorf("Test failed: %s\nExpected: %#v\nGot: %#v", tt.name, tt.expected, expression)
			}
		})
	}
}

func TestParser_InvalidCode(t *testing.T) {
	tests := []struct {
		name        string
		tokens      []token.Token
		expectedErr string
	}{
		{
			name: "Mismatched Parentheses (1 + 2",
			tokens: []token.Token{
				{Type: token.LEFT_PAREN, Literal: "("},
				{Type: token.NUMBER, Literal: 1},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.NUMBER, Literal: 2},
				{Type: token.EOF},
			},
			expectedErr: "Expect ')' after expression.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if err, ok := r.(*error.Error); ok {
						if err.Message != tt.expectedErr {
							t.Errorf("Expected error message '%s' but got '%s'", tt.expectedErr, err.Message)
						}
					} else {
						t.Errorf("Expected a parse error but got %v", r)
					}
				} else {
					t.Errorf("Expected an error but no error was raised")
				}
			}()

			// Run the parser, expecting an error to be raised
			p := New(tt.tokens)

			p.expression() // Should trigger error
		})
	}
}
