package parser

import (
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

func ProcessReturnStatement(t *testing.T, s statements.Statement, value string) bool {
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
