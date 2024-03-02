package parser

import (
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
