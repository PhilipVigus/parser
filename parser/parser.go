package parser

import (
	"lang/ast/expressions"
	"lang/ast/statements"
	"lang/lexer"
	"lang/lexer/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *statements.Program {
	program := &statements.Program{}
	program.Statements = []statements.Statement{}
	for !p.curTokenIs(token.Eof) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() statements.Statement {
	switch p.currentToken.Type {
	case token.Let:
		return p.parseAssignStatement()
	default:
		return nil
	}
}

func (p *Parser) parseAssignStatement() *statements.Assign {
	stmt := &statements.Assign{Token: p.currentToken}

	if !p.expectPeek(token.Ident) {
		return nil
	}

	stmt.Name = &expressions.Identifier{Token: p.currentToken, Value: p.currentToken.Value}

	if !p.expectPeek(token.Assign) {
		return nil
	}

	for !p.curTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
