package parser

import (
	"fmt"
	"lang/ast/expressions"
	"lang/ast/statements"
	"lang/lexer"
	"lang/lexer/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	p.errors = []string{}
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		token.GetStringFromTokenType(t), token.GetStringFromTokenType(p.peekToken.Type))
	p.errors = append(p.errors, msg)
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
	case token.Return:
		return p.parseReturnStatement()
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

func (p *Parser) parseReturnStatement() *statements.Return {
	stmt := &statements.Return{Token: p.currentToken}

	p.nextToken()

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
		p.peekError(t)
		return false
	}
}
