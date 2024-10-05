/*
Package printer provides a visitor that prints the AST
*/
package printer

import (
	"fmt"
	"golox/expr"
	"strings"
)

// AstPrinter is a visitor that prints the AST
type AstPrinter struct{}

// New creates a new AstPrinter
func New() *AstPrinter {
	return &AstPrinter{}
}

// Print the expression
func (a *AstPrinter) Print(e expr.Expr) string {
	return e.Accept(a).(string)
}

// VisitBinaryExpr implements the Visitor interface
func (a *AstPrinter) VisitBinaryExpr(e *expr.Binary) interface{} {
	return a.parenthesize(e.Operator.Lexeme, e.Left, e.Right)
}

// VisitGroupingExpr implements the Visitor interface
func (a *AstPrinter) VisitGroupingExpr(e *expr.Grouping) interface{} {
	return a.parenthesize("group", e.Expression)
}

// VisitLiteralExpr implements the Visitor interface
func (a *AstPrinter) VisitLiteralExpr(e *expr.Literal) interface{} {
	if e.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", e.Value)
}

// VisitUnaryExpr implements the Visitor interface
func (a *AstPrinter) VisitUnaryExpr(e *expr.Unary) interface{} {
	return a.parenthesize(e.Operator.Lexeme, e.Right)
}

// VisitVariableExpr implements the Visitor interface
func (a *AstPrinter) VisitVariableExpr(e *expr.Variable) interface{} {
	return e.Name.Lexeme
}

// VisitAssignExpr implements the Visitor interface
func (a *AstPrinter) VisitAssignExpr(e *expr.Assign) interface{} {
	return a.parenthesize("=", e.Name.Lexeme, e.Value)
}

// VisitLogicalExpr implements the Visitor interface
func (a *AstPrinter) VisitLogicalExpr(e *expr.Logical) interface{} {
	return a.parenthesize(e.Operator.Lexeme, e.Left, e.Right)
}

// VisitCallExpr implements the Visitor interface
func (a *AstPrinter) VisitCallExpr(call *expr.Call) interface{} {
	var args strings.Builder

	for idx, ele := range call.Arguments {
		args.WriteString(ele.Accept(a).(string))
		if idx != len(call.Arguments)-1 {
			args.WriteString(",")
		}
	}

	if args.Len() > 0 {
		return a.parenthesize("call", call.Callee, args.String())
	}

	return a.parenthesize("call", call.Callee)
}

// VisitGetExpr implements the Visitor interface
func (a *AstPrinter) VisitGetExpr(e *expr.Get) interface{} {
	return a.parenthesize("get", e.Object, e.Name.Lexeme)
}

// VisitSetExpr implements the Visitor interface
func (a *AstPrinter) VisitSetExpr(e *expr.Set) interface{} {
	return a.parenthesize("set", e.Object, e.Name.Lexeme, e.Value)
}

// VisitThisExpr implements the Visitor interface
func (a *AstPrinter) VisitThisExpr(_ *expr.This) interface{} {
	return "this"
}

// VisitSuperExpr implements the Visitor interface
func (a *AstPrinter) VisitSuperExpr(_ *expr.Super) interface{} {
	return "super"
}

func (a *AstPrinter) parenthesize(name string, parts ...interface{}) string {
	var str strings.Builder

	str.WriteString("(")
	str.WriteString(name)

	for _, part := range parts {
		str.WriteString(" ")
		switch p := part.(type) {
		case expr.Expr:
			str.WriteString(p.Accept(a).(string))
		case string:
			str.WriteString(p)
		case fmt.Stringer:
			str.WriteString(p.String())
		}
	}
	str.WriteString(")")

	return str.String()
}
