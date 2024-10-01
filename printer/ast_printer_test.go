package printer

import (
	"golox/expr"
	"golox/token"
	"testing"
)

func TestAstPrinter_SingleExprs(t *testing.T) {
	tests := []struct {
		name     string
		expr     expr.Expr
		expected string
	}{
		{
			name: "Literal expression - number",
			expr: &expr.Literal{
				Value: 123,
			},
			expected: "123",
		},
		{
			name: "Literal expression - string",
			expr: &expr.Literal{
				Value: "hello",
			},
			expected: "hello",
		},
		{
			name: "Unary expression",
			expr: &expr.Unary{
				Operator: &token.Token{
					Lexeme: "-",
				},
				Right: &expr.Literal{
					Value: 123,
				},
			},
			expected: "(- 123)",
		},
		{
			name: "Binary expression",
			expr: &expr.Binary{
				Left: &expr.Literal{
					Value: 1,
				},
				Operator: &token.Token{
					Lexeme: "+",
				},
				Right: &expr.Literal{
					Value: 2,
				},
			},
			expected: "(+ 1 2)",
		},
		{
			name: "Grouping expression",
			expr: &expr.Grouping{
				Expression: &expr.Literal{
					Value: 45.67,
				},
			},
			expected: "(group 45.67)",
		},
		{
			name: "Variable expression",
			expr: &expr.Variable{
				Name: &token.Token{
					Lexeme: "x",
				},
			},
			expected: "x",
		},
		{
			name: "Assign expression",
			expr: &expr.Assign{
				Name: &token.Token{
					Lexeme: "x",
				},
				Value: &expr.Literal{
					Value: 42,
				},
			},
			expected: "(= x 42)",
		},
		{
			name: "Logical expression",
			expr: &expr.Logical{
				Left: &expr.Literal{
					Value: true,
				},
				Operator: &token.Token{
					Lexeme: "or",
				},
				Right: &expr.Literal{
					Value: false,
				},
			},
			expected: "(or true false)",
		},
		{
			name: "Call expression",
			expr: &expr.Call{
				Callee: &expr.Variable{
					Name: &token.Token{
						Lexeme: "func",
					},
				},
				Arguments: []expr.Expr{
					&expr.Literal{Value: 1},
					&expr.Literal{Value: 2},
				},
			},
			expected: "(call func 1,2)",
		},
		{
			name: "Get expression",
			expr: &expr.Get{
				Object: &expr.Variable{
					Name: &token.Token{
						Lexeme: "object",
					},
				},
				Name: &token.Token{
					Lexeme: "property",
				},
			},
			expected: "(get object property)",
		},
		{
			name: "Set expression",
			expr: &expr.Set{
				Object: &expr.Variable{
					Name: &token.Token{
						Lexeme: "object",
					},
				},
				Name: &token.Token{
					Lexeme: "property",
				},
				Value: &expr.Literal{
					Value: 3,
				},
			},
			expected: "(set object property 3)",
		},
		{
			name:     "This expression",
			expr:     &expr.This{},
			expected: "this",
		},
		{
			name:     "Super expression",
			expr:     &expr.Super{},
			expected: "super",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := PrintAst(tt.expr)

			if actual != tt.expected {
				t.Errorf("PrintAst() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestAstPrinter_ComplexExprs(t *testing.T) {
	tests := []struct {
		name     string
		expr     expr.Expr
		expected string
	}{
		{
			name: "Complex expression 1: (1 + 2) * (3 - 4)",
			expr: &expr.Binary{
				Left: &expr.Grouping{
					Expression: &expr.Binary{
						Left: &expr.Literal{Value: 1},
						Operator: &token.Token{
							Lexeme: "+",
						},
						Right: &expr.Literal{Value: 2},
					},
				},
				Operator: &token.Token{
					Lexeme: "*",
				},
				Right: &expr.Grouping{
					Expression: &expr.Binary{
						Left: &expr.Literal{Value: 3},
						Operator: &token.Token{
							Lexeme: "-",
						},
						Right: &expr.Literal{Value: 4},
					},
				},
			},
			expected: "(* (group (+ 1 2)) (group (- 3 4)))",
		},
		{
			name: "Complex expression 2: -((5 + 6) / 7)",
			expr: &expr.Unary{
				Operator: &token.Token{
					Lexeme: "-",
				},
				Right: &expr.Grouping{
					Expression: &expr.Binary{
						Left: &expr.Binary{
							Left: &expr.Literal{Value: 5},
							Operator: &token.Token{
								Lexeme: "+",
							},
							Right: &expr.Literal{Value: 6},
						},
						Operator: &token.Token{
							Lexeme: "/",
						},
						Right: &expr.Literal{Value: 7},
					},
				},
			},
			expected: "(- (group (/ (+ 5 6) 7)))",
		},
		{
			name: "Complex expression 3: 10 / (3 * (4 + 5))",
			expr: &expr.Binary{
				Left: &expr.Literal{
					Value: 10,
				},
				Operator: &token.Token{
					Lexeme: "/",
				},
				Right: &expr.Grouping{
					Expression: &expr.Binary{
						Left: &expr.Literal{
							Value: 3,
						},
						Operator: &token.Token{
							Lexeme: "*",
						},
						Right: &expr.Grouping{
							Expression: &expr.Binary{
								Left: &expr.Literal{
									Value: 4,
								},
								Operator: &token.Token{
									Lexeme: "+",
								},
								Right: &expr.Literal{
									Value: 5,
								},
							},
						},
					},
				},
			},
			expected: "(/ 10 (group (* 3 (group (+ 4 5)))))",
		},
		{
			name: "Complex expression 4: -(3 * (1 + 2)) + 5",
			expr: &expr.Binary{
				Left: &expr.Unary{
					Operator: &token.Token{
						Lexeme: "-",
					},
					Right: &expr.Grouping{
						Expression: &expr.Binary{
							Left: &expr.Literal{
								Value: 3,
							},
							Operator: &token.Token{
								Lexeme: "*",
							},
							Right: &expr.Grouping{
								Expression: &expr.Binary{
									Left: &expr.Literal{
										Value: 1,
									},
									Operator: &token.Token{
										Lexeme: "+",
									},
									Right: &expr.Literal{
										Value: 2,
									},
								},
							},
						},
					},
				},
				Operator: &token.Token{
					Lexeme: "+",
				},
				Right: &expr.Literal{
					Value: 5,
				},
			},
			expected: "(+ (- (group (* 3 (group (+ 1 2))))) 5)",
		},
		{
			name: "Complex expression 5: (7 - (4 / (2 + 1))) * 3",
			expr: &expr.Binary{
				Left: &expr.Grouping{
					Expression: &expr.Binary{
						Left: &expr.Literal{Value: 7},
						Operator: &token.Token{
							Lexeme: "-",
						},
						Right: &expr.Grouping{
							Expression: &expr.Binary{
								Left: &expr.Literal{
									Value: 4,
								},
								Operator: &token.Token{
									Lexeme: "/",
								},
								Right: &expr.Grouping{
									Expression: &expr.Binary{
										Left: &expr.Literal{Value: 2},
										Operator: &token.Token{
											Lexeme: "+",
										},
										Right: &expr.Literal{Value: 1},
									},
								},
							},
						},
					},
				},
				Operator: &token.Token{
					Lexeme: "*",
				},
				Right: &expr.Literal{Value: 3},
			},
			expected: "(* (group (- 7 (group (/ 4 (group (+ 2 1)))))) 3)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := PrintAst(tt.expr)

			if actual != tt.expected {
				t.Errorf("PrintAst() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
