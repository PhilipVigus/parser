package parser

import (
	"fmt"
	"lang/ast/expressions"
	"lang/ast/statements"
	"lang/lexer"
	"strings"
	"testing"
)

func TestAssignmentStatements(t *testing.T) {
	input := `
		let x = 5;
		let foo = 10;`

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"foo"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !ProcessAssignmentStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 10;`

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"5"},
		{"10"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !ProcessReturnStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func ProcessAssignmentStatement(t *testing.T, s statements.Statement, name string) bool {
	if s.TokenValue() != "let" {
		t.Errorf("s.TokenValue not 'let'. got=%q", s.TokenValue())
		return false
	}

	assignmentStatement, ok := s.(*statements.Assign)
	if !ok {
		t.Errorf("s not *ast.AssignmentStatement. got=%T", s)
		return false
	}

	if assignmentStatement.Name.Value != name {
		t.Errorf("AssignmentStatement.Name.Value not '%s'. got=%s", name, assignmentStatement.Name.Value)
		return false
	}

	if assignmentStatement.Name.TokenValue() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, assignmentStatement.Name)
		return false
	}

	return true
}

func ProcessReturnStatement(t *testing.T, s statements.Statement, _ string) bool {
	if s.TokenValue() != "return" {
		t.Errorf("s.TokenValue not 'return'. got=%q", s.TokenValue())
		return false
	}

	_, ok := s.(*statements.Return)
	if !ok {
		t.Errorf("s not *ast.ReturnStatement. got=%T", s)
		return false
	}

	return true
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()

	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*statements.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*expressions.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}

	if ident.TokenValue() != "foobar" {
		t.Errorf("ident.TokenValue not %s. got=%s", "foobar", ident.TokenValue())
	}
}

func TestNumberLiteralExpression(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
		expectedType  interface{}
	}{
		{"5;", 5, int64(0)},       // Example integer test case
		{"5.5;", 5.5, float64(0)}, // Example float test case
		// Add more test cases as needed
	}

	for _, tt := range tests {
		l := lexer.New(strings.NewReader(tt.input))
		p := New(l)
		program := p.ParseProgram()

		checkParseErrors(t, p) // Ensure your test setup includes this function to check parser errors

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*statements.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not *statements.ExpressionStatement. got=%T", program.Statements[0])
		}

		switch expected := tt.expectedValue.(type) {
		case int:
			literal, ok := stmt.Expression.(*expressions.NumberLiteral[int64])
			if !ok {
				t.Errorf("for input '%s', expected *expressions.NumberLiteral[int64], got=%T", tt.input, stmt.Expression)
			} else if int64(expected) != literal.Value {
				t.Errorf("for input '%s', expected value %d, got=%d", tt.input, expected, literal.Value)
			}
		case float64:
			literal, ok := stmt.Expression.(*expressions.NumberLiteral[float64])
			if !ok {
				t.Errorf("for input '%s', expected *expressions.NumberLiteral[float64], got=%T", tt.input, stmt.Expression)
			} else if expected != literal.Value {
				t.Errorf("for input '%s', expected value %f, got=%f", tt.input, expected, literal.Value)
			}
		default:
			t.Fatalf("unsupported type in test case")
		}

		if tt.expectedType != nil && stmt.Expression.String() != tt.input[:len(tt.input)-1] {
			t.Errorf("literal.TokenValue not %s. got=%s", tt.input[:len(tt.input)-1], stmt.Expression.String())
		}
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!5;", "!", 5},
		{"-15.5;", "-", 15.5},
		// Add more test cases as needed
	}

	for _, tt := range prefixTests {
		l := lexer.New(strings.NewReader(tt.input))
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*statements.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*expressions.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}

		switch expected := tt.value.(type) {
		case int:
			if !IntegerLiteralExpressionTester(t, exp.Right, int64(expected)) {
				return
			}
		case float64:
			if !FloatLiteralExpressionTester(t, exp.Right, expected) {
				return
			}
		default:
			t.Fatalf("unsupported type in test case")
		}
	}
}

func IntegerLiteralExpressionTester(t *testing.T, il expressions.Expression, value int64) bool {
	i, ok := il.(*expressions.NumberLiteral[int64])
	if !ok {
		t.Errorf("il not IntegerLiteral. got=%T", il)
		return false
	}

	if i.Value != value {
		t.Errorf("i.Value not %d. got=%d", value, i.Value)
		return false
	}

	if i.TokenValue() != fmt.Sprintf("%d", value) {
		t.Errorf("i.TokenLiteral not %d. got=%s", value,
			i.TokenValue())
		return false
	}

	return true
}

func FloatLiteralExpressionTester(t *testing.T, il expressions.Expression, value float64) bool {
	f, ok := il.(*expressions.NumberLiteral[float64])
	if !ok {
		t.Errorf("f not IntegerLiteral. got=%T", il)
		return false
	}

	f.TokenValue()

	if f.Value != value {
		t.Errorf("f.Value not %f. got=%f", value, f.Value)
		return false
	}

	if !CompareFloatingPointStrings(f.TokenValue(), fmt.Sprintf("%f", value)) {
		t.Errorf("f.TokenLiteral not %f. got=%s", value, f.TokenValue())
		return false
	}

	return true
}

// CompareFloatingPointStrings compares two strings representing floating point numbers.
// It trims trailing zeros and decimal points if they result in an integer value.
func CompareFloatingPointStrings(a, b string) bool {
	// Trim trailing zeros and decimal points if they result in an integer value.
	trimmedA := strings.TrimRight(strings.TrimRight(a, "0"), ".")
	trimmedB := strings.TrimRight(strings.TrimRight(b, "0"), ".")

	return trimmedA == trimmedB
}
