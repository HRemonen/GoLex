package printer

import (
	"golox/expr"
	"golox/token"
	"testing"
)

func TestAstPrinter_Print(t *testing.T) {
	minus := token.Token{Type: token.MINUS, Lexeme: "-"}
	star := token.Token{Type: token.STAR, Lexeme: "*"}
	grouping := expr.Grouping{Expression: &expr.Literal{Value: 45.67}}
	unary := expr.Unary{Operator: &minus, Right: &expr.Literal{Value: 123}}
	binary := expr.Binary{Left: &unary, Operator: &star, Right: &grouping}

	str := PrintAst(&binary)

	expected := "(* (- 123) (group 45.67))"
	if str != expected {
		t.Errorf("AstPrinter.Print() = %v, want %v", str, expected)
	}
}
