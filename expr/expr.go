/*
Package expr contains the AST nodes for the expressions in the Lox language.
*/
package expr

import (
	"golox/token"
)

// Expr is the interface that all expressions must implement
type Expr interface {
	Accept(v Visitor) interface{}
}

// Visitor is the interface that all visitors must implement
type Visitor interface {
	VisitAssignExpr(expr *Assign) interface{}
	VisitBinaryExpr(expr *Binary) interface{}
	VisitCallExpr(expr *Call) interface{}
	VisitGetExpr(expr *Get) interface{}
	VisitGroupingExpr(expr *Grouping) interface{}
	VisitLiteralExpr(expr *Literal) interface{}
	VisitLogicalExpr(expr *Logical) interface{}
	VisitSetExpr(expr *Set) interface{}
	VisitSuperExpr(expr *Super) interface{}
	VisitThisExpr(expr *This) interface{}
	VisitUnaryExpr(expr *Unary) interface{}
	VisitVariableExpr(expr *Variable) interface{}
}

// Assign represents an assignment expression
type Assign struct {
	Name  *token.Token
	Value Expr
}

// Accept implements the Expr interface
func (e *Assign) Accept(v Visitor) interface{} {
	return v.VisitAssignExpr(e)
}

// Binary represents a binary expression
type Binary struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

// Accept implements the Expr interface
func (e *Binary) Accept(v Visitor) interface{} {
	return v.VisitBinaryExpr(e)
}

// Call represents a call expression
type Call struct {
	Callee    Expr
	Paren     *token.Token
	Arguments []Expr
}

// Accept implements the Expr interface
func (e *Call) Accept(v Visitor) interface{} {
	return v.VisitCallExpr(e)
}

// Get represents a get expression
type Get struct {
	Object Expr
	Name   *token.Token
}

// Accept implements the Expr interface
func (e *Get) Accept(v Visitor) interface{} {
	return v.VisitGetExpr(e)
}

// Grouping represents a grouping expression
type Grouping struct {
	Expression Expr
}

// Accept implements the Expr interface
func (e *Grouping) Accept(v Visitor) interface{} {
	return v.VisitGroupingExpr(e)
}

// Literal represents a literal expression
type Literal struct {
	Value interface{}
}

// Accept implements the Expr interface
func (e *Literal) Accept(v Visitor) interface{} {
	return v.VisitLiteralExpr(e)
}

// Logical represents a logical expression
type Logical struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

// Accept implements the Expr interface
func (e *Logical) Accept(v Visitor) interface{} {
	return v.VisitLogicalExpr(e)
}

// Set represents a set expression
type Set struct {
	Object Expr
	Name   *token.Token
	Value  Expr
}

// Accept implements the Expr interface
func (e *Set) Accept(v Visitor) interface{} {
	return v.VisitSetExpr(e)
}

// Super represents a super expression
type Super struct {
	Keyword *token.Token
	Method  *token.Token
}

// Accept implements the Expr interface
func (e *Super) Accept(v Visitor) interface{} {
	return v.VisitSuperExpr(e)
}

// This represents a this expression
type This struct {
	Keyword *token.Token
}

// Accept implements the Expr interface
func (e *This) Accept(v Visitor) interface{} {
	return v.VisitThisExpr(e)
}

// Unary represents a unary expression
type Unary struct {
	Operator *token.Token
	Right    Expr
}

// Accept implements the Expr interface
func (e *Unary) Accept(v Visitor) interface{} {
	return v.VisitUnaryExpr(e)
}

// Variable represents a variable expression
type Variable struct {
	Name *token.Token
}

// Accept implements the Expr interface
func (e *Variable) Accept(v Visitor) interface{} {
	return v.VisitVariableExpr(e)
}
