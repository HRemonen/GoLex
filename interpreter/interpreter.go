package interpreter

import (
	"golox/expr"
	"golox/token"
)

// Interpreter is the visitor that interprets the AST
type Interpreter struct{}

// New creates a new Interpreter
func New() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) VisitLiteralExpr(e *expr.Literal) interface{} {
	return e.Value
}

func (i *Interpreter) VisitGroupingExpr(e *expr.Grouping) interface{} {
	return i.evaluate(e.Expression)
}

func (i *Interpreter) VisitUnaryExpr(e *expr.Unary) interface{} {
	right := i.evaluate(e.Right)

	switch e.Operator.Type {
	case token.BANG:
		return !isTruthy(right)
	case token.MINUS:
		checkNumberOperand(e.Operator, right)
		
		return -right.(float64)
	}

	// Unreachable
	return nil
}

func (i *Interpreter) VisitBinaryExpr(e *expr.Binary) interface{} {
	left := i.evaluate(e.Left)
	right := i.evaluate(e.Right)

	switch e.Operator.Type {
	case token.GREATER:
		return left.(float64) > right.(float64)
	case token.GREATER_EQUAL:
		return left.(float64) >= right.(float64)
	case token.LESS:
		return left.(float64) < right.(float64)
	case token.LESS_EQUAL:
		return left.(float64) <= right.(float64)
	case token.BANG_EQUAL:
		return !isEqual(left, right)
	case token.EQUAL_EQUAL:
		return isEqual(left, right)
	case token.MINUS:
		return left.(float64) - right.(float64)
	case token.PLUS:
		if l, ok := left.(float64); ok {
			if r, ok := right.(float64); ok {
				return l + r
			}
		}

		if l, ok := left.(string); ok {
			if r, ok := right.(string); ok {
				return l + r
			}
		}
	case token.SLASH:
		return left.(float64) / right.(float64)
	case token.STAR:
		return left.(float64) * right.(float64)
	}

	// Unreachable
	return nil
}

func (i *Interpreter) evaluate(e expr.Expr) interface{} {
	return e.Accept(i)
}

// We follow simple rule to determine truthiness:
// - nil and false are false
// - everything else is true
func isTruthy(value interface{}) bool {
	switch v := value.(type) {
	case nil:
		return false
	case bool:
		return v
	default:
		return true
	}
}

func isEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil {
		return false
	}

	return a == b
}

func checkNumberOperand(operator *token.Token, operand interface{}) {
	if _, ok := operand.(float64); !ok {
		panic("Invalid operation: operator '" + operator.Lexeme + "' not defined on '" + operand.(string) + "'")
	}
}
