/*
Package stmt defines the statements that can be used in the Lox language.
*/
package stmt

import (
	"golox/expr"
	"golox/token"
)

// Stmt is the interface that all statements must implement
type Stmt interface {
	Accept(v Visitor) interface{}
}

// Visitor is the interface that all visitors must implement
type Visitor interface {
	VisitBlockStmt(stmt *Block) interface{}
	VisitClassStmt(stmt *Class) interface{}
	VisitExpressionStmt(stmt *Expression) interface{}
	VisitFunctionStmt(stmt *Function) interface{}
	VisitIfStmt(stmt *If) interface{}
	VisitPrintStmt(stmt *Print) interface{}
	VisitReturnStmt(stmt *Return) interface{}
	VisitVarStmt(stmt *Var) interface{}
	VisitWhileStmt(stmt *While) interface{}
}

// Block represents a block statement
type Block struct {
	Statements []Stmt
}

// Accept implements the Stmt interface
func (s *Block) Accept(v Visitor) interface{} {
	return v.VisitBlockStmt(s)
}

// Class represents a class statement
type Class struct {
	Name       *token.Token
	Superclass *expr.Variable
	Methods    []*Function
}

// Accept implements the Stmt interface
func (s *Class) Accept(v Visitor) interface{} {
	return v.VisitClassStmt(s)
}

// Expression represents an expression statement
type Expression struct {
	Expression expr.Expr
}

// Accept implements the Stmt interface
func (s *Expression) Accept(v Visitor) interface{} {
	return v.VisitExpressionStmt(s)
}

// Function represents a function statement
type Function struct {
	Name   *token.Token
	Params []*token.Token
	Body   []Stmt
}

// Accept implements the Stmt interface
func (s *Function) Accept(v Visitor) interface{} {
	return v.VisitFunctionStmt(s)
}

// If represents an if statement
type If struct {
	Condition  expr.Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

// Accept implements the Stmt interface
func (s *If) Accept(v Visitor) interface{} {
	return v.VisitIfStmt(s)
}

// Print represents a print statement
type Print struct {
	Expression expr.Expr
}

// Accept implements the Stmt interface
func (s *Print) Accept(v Visitor) interface{} {
	return v.VisitPrintStmt(s)
}

// Return represents a return statement
type Return struct {
	Keyword *token.Token
	Value   expr.Expr
}

// Accept implements the Stmt interface
func (s *Return) Accept(v Visitor) interface{} {
	return v.VisitReturnStmt(s)
}

// Var represents a var statement
type Var struct {
	Name        *token.Token
	Initializer expr.Expr
}

// Accept implements the Stmt interface
func (s *Var) Accept(v Visitor) interface{} {
	return v.VisitVarStmt(s)
}

// While represents a while statement
type While struct {
	Condition expr.Expr
	Body      Stmt
}

// Accept implements the Stmt interface
func (s *While) Accept(v Visitor) interface{} {
	return v.VisitWhileStmt(s)
}
