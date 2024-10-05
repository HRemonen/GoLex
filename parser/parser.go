package parser

import (
	"golox/error"
	"golox/expr"
	"golox/token"
)

type Parser struct {
	tokens  []token.Token
	current int // Next token to be parsed
}

func New(tokens []token.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

// Expression maps to the CFG rule: expression → equality ;
func (p *Parser) expression() expr.Expr {
	return p.equality()
}

// Equality maps to the CFG rule: equality → comparison ( ( "!=" | "==" ) comparison )* ;
// comparison is the first non-terminal in the rule
// ( ( "!=" | "==" ) comparison )* is the optional part of the rule
//
// The rule is left-associative
// Grabs the matched operator token and the right operand and creates a new Binary expression
// with the left operand, operator and right operand
//
// The loop continues until there are no more matched operators
func (p *Parser) equality() expr.Expr {
	// first comparison non-terminal in the rule
	expression := p.comparison()

	// loop through the optional ( ( "!=" | "==" ) comparison )* part of the rule
	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expression = &expr.Binary{Left: expression, Operator: operator, Right: right}
	}

	return expression
}

// Comparison maps to the CFG rule: comparison → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
// term is the first non-terminal in the rule
// ( ( ">" | ">=" | "<" | "<=" ) term )* is the optional part of the rule
//
// The rule is left-associative
// Grabs the matched operator token and the right operand and creates a new Binary expression
// with the left operand, operator and right operand
//
// The loop continues until there are no more matched operators
func (p *Parser) comparison() expr.Expr {
	// first term non-terminal in the rule
	expression := p.term()

	// loop through the optional ( ( ">" | ">=" | "<" | "<=" ) term )* part of the rule
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expression = &expr.Binary{Left: expression, Operator: operator, Right: right}
	}

	return expression
}

// Term maps to the CFG rule: term → factor ( ( "-" | "+" ) factor )* ;
func (p *Parser) term() expr.Expr {
	expression := p.factor()

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right := p.factor()
		expression = &expr.Binary{Left: expression, Operator: operator, Right: right}
	}

	return expression
}

// Factor maps to the CFG rule: factor → unary ( ( "/" | "*" ) unary )* ;
func (p *Parser) factor() expr.Expr {
	expression := p.unary()

	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right := p.unary()
		expression = &expr.Binary{Left: expression, Operator: operator, Right: right}
	}

	return expression
}

// Unary maps to the CFG rule: unary → ( "!" | "-" ) unary | primary ;
func (p *Parser) unary() expr.Expr {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.unary()
		return &expr.Unary{Operator: operator, Right: right}
	}

	return p.primary()
}

// Primary maps to the CFG rule: primary → NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")" ;
func (p *Parser) primary() expr.Expr {
	switch {
	case p.match(token.FALSE):
		return &expr.Literal{Value: false}
	case p.match(token.TRUE):
		return &expr.Literal{Value: true}
	case p.match(token.NULL):
		return &expr.Literal{Value: nil}
	case p.match(token.NUMBER, token.STRING):
		return &expr.Literal{Value: p.previous().Literal}
	case p.match(token.LEFT_PAREN):
		expression := p.expression()
		p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		return &expr.Grouping{Expression: expression}
	}

	return nil
}

// Check if the current token is any of the given types. If it does, consume it
func (p *Parser) match(types ...token.Type) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

// Consume the current token if it is of the given type. If it is not, panic with the given message
func (p *Parser) consume(t token.Type, message string) *token.Token {
	if p.check(t) {
		return p.advance()
	}

	if err := parseError(p.peek(), message); err != nil {
		panic(err)
	}

	return nil
}

// Check if the current token is of the given type without consuming it
func (p *Parser) check(t token.Type) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

// Consume the current token and return it
func (p *Parser) advance() *token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

// Check if we have reached the end of the token list
func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

// Return the current token yet to be consumed
func (p *Parser) peek() *token.Token {
	return &p.tokens[p.current]
}

// Return the previous token that was consumed
func (p *Parser) previous() *token.Token {
	return &p.tokens[p.current-1]
}

func parseError(t *token.Token, message string) *error.Error {
	return error.New(t, message)
}

// Synchronize the parser after an error has been encountered
// This is done by skipping tokens until a statement boundary is reached
func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == token.SEMICOLON {
			return
		}

		switch p.peek().Type {
		case token.CLASS, token.FUN, token.VAR, token.FOR, token.IF, token.WHILE, token.PRINT, token.RETURN:
			return
		}

		p.advance()
	}
}
