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
		{
			name: "Ternary Operator (true ? 1 : 2)",
			tokens: []token.Token{
				{Type: token.TRUE, Literal: true},
				{Type: token.QUESTION, Literal: "?"},
				{Type: token.NUMBER, Literal: 1},
				{Type: token.COLON, Literal: ":"},
				{Type: token.NUMBER, Literal: 2},
				{Type: token.EOF},
			},
			expected: &expr.Ternary{
				Condition:   &expr.Literal{Value: true},
				TrueBranch:  &expr.Literal{Value: 1},
				FalseBranch: &expr.Literal{Value: 2},
			},
		},
		{
			name: "Ternary Operator with Nested Ternary (true ? 1 : false ? 2 : 3)",
			tokens: []token.Token{
				{Type: token.TRUE, Literal: true},
				{Type: token.QUESTION, Literal: "?"},
				{Type: token.NUMBER, Literal: 1},
				{Type: token.COLON, Literal: ":"},
				{Type: token.FALSE, Literal: false},
				{Type: token.QUESTION, Literal: "?"},
				{Type: token.NUMBER, Literal: 2},
				{Type: token.COLON, Literal: ":"},
				{Type: token.NUMBER, Literal: 3},
				{Type: token.EOF},
			},
			expected: &expr.Ternary{
				Condition:  &expr.Literal{Value: true},
				TrueBranch: &expr.Literal{Value: 1},
				FalseBranch: &expr.Ternary{
					Condition:   &expr.Literal{Value: false},
					TrueBranch:  &expr.Literal{Value: 2},
					FalseBranch: &expr.Literal{Value: 3},
				},
			},
		},
		{
			name: "Ternary Operator with Nested Ternary (false ? true ? 1 : 2 : 3)",
			tokens: []token.Token{
				{Type: token.FALSE, Literal: false},
				{Type: token.QUESTION, Literal: "?"},
				{Type: token.TRUE, Literal: true},
				{Type: token.QUESTION, Literal: "?"},
				{Type: token.NUMBER, Literal: 1},
				{Type: token.COLON, Literal: ":"},
				{Type: token.NUMBER, Literal: 2},
				{Type: token.COLON, Literal: ":"},
				{Type: token.NUMBER, Literal: 3},
				{Type: token.EOF},
			},
			expected: &expr.Ternary{
				Condition: &expr.Literal{Value: false},
				TrueBranch: &expr.Ternary{
					Condition:   &expr.Literal{Value: true},
					TrueBranch:  &expr.Literal{Value: 1},
					FalseBranch: &expr.Literal{Value: 2},
				},
				FalseBranch: &expr.Literal{Value: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.tokens)

			expression := p.Parse()

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
		{
			name: "Missing expression in parentheses (1 + )",
			tokens: []token.Token{
				{Type: token.LEFT_PAREN, Literal: "("},
				{Type: token.NUMBER, Literal: 1},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.RIGHT_PAREN, Literal: ")"},
				{Type: token.EOF},
			},
			expectedErr: "Expect expression.",
		},
		{
			name: "Empty expression",
			tokens: []token.Token{
				{Type: token.EOF},
			},
			expectedErr: "Expect expression.",
		},
		{
			name: "Missing false ternary branch (true ? 1 )",
			tokens: []token.Token{
				{Type: token.TRUE, Literal: true},
				{Type: token.QUESTION, Literal: "?"},
				{Type: token.NUMBER, Literal: 1},
				{Type: token.EOF},
			},
			expectedErr: "Expect ':' after true branch of ternary expression.",
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

			p := New(tt.tokens)

			p.Parse()
		})
	}
}
