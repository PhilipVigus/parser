package parser

import (
	"fmt"
	"lang/ast/expressions"
	"lang/ast/statements"
	"lang/lexer"
	"lang/lexer/token"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS       // ==
	LESS_GREATER // > or <
	SUM          // +
	PRODUCT      // *
	PREFIX       // -X or !X
	CALL         // myFunction(X)
)

type (
	prefixParseFn func() expressions.Expression
	infixParseFn  func(expressions.Expression) expressions.Expression
)

type Parser struct {
	l              *lexer.Lexer
	currentToken   token.Token
	peekToken      token.Token
	errors         []string
	prefixParseFns map[token.Type]prefixParseFn
	infixParseFns  map[token.Type]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	p.errors = []string{}

	p.prefixParseFns = make(map[token.Type]prefixParseFn)
	p.infixParseFns = make(map[token.Type]infixParseFn)

	p.registerPrefix(token.Ident, p.parseIdentifier)
	p.registerPrefix(token.Number, p.ParseNumberLiteral)
	p.registerPrefix(token.Minus, p.parsePrefixExpression)
	p.registerPrefix(token.Not, p.parsePrefixExpression)
	p.registerInfix(token.Plus, p.parseInfixExpression)
	p.registerInfix(token.Minus, p.parseInfixExpression)
	p.registerInfix(token.Multiply, p.parseInfixExpression)
	p.registerInfix(token.Divide, p.parseInfixExpression)
	p.registerInfix(token.Equal, p.parseInfixExpression)
	p.registerInfix(token.NotEqual, p.parseInfixExpression)
	p.registerInfix(token.LessThan, p.parseInfixExpression)
	p.registerInfix(token.GreaterThan, p.parseInfixExpression)
	p.registerInfix(token.LessThanOrEqual, p.parseInfixExpression)
	p.registerInfix(token.GreaterThanOrEqual, p.parseInfixExpression)

	return p
}

func (p *Parser) registerPrefix(t token.Type, fn prefixParseFn) {
	p.prefixParseFns[t] = fn
}

func (p *Parser) registerInfix(t token.Type, fn infixParseFn) {
	p.infixParseFns[t] = fn
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
		return p.parseExpressionStatement()
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

func (p *Parser) parseExpressionStatement() *statements.ExpressionStatement {
	stmt := &statements.ExpressionStatement{Token: p.currentToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) expressions.Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}
	leftExp := prefix()
	for !p.peekTokenIs(token.Semicolon) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) parseIdentifier() expressions.Expression {
	return &expressions.Identifier{Token: p.currentToken, Value: p.currentToken.Value}
}

func (p *Parser) ParseNumberLiteral() expressions.Expression {
	tokenValue := p.currentToken.Value

	if i, err := strconv.ParseInt(tokenValue, 10, 64); err == nil {
		return &expressions.NumberLiteral[int64]{
			Token: p.currentToken,
			Value: i,
		}
	}

	if f, err := strconv.ParseFloat(tokenValue, 64); err == nil {
		return &expressions.NumberLiteral[float64]{
			Token: p.currentToken,
			Value: f,
		}
	}

	p.errors = append(p.errors, fmt.Sprintf("could not parse %q as number", tokenValue))
	return nil
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

func (p *Parser) noPrefixParseFnError(t token.Type) {
	msg := fmt.Sprintf("no prefix parse function for %s found", token.GetStringFromTokenType(t))
	p.errors = append(p.errors, msg)
}

func (p *Parser) parsePrefixExpression() expressions.Expression {
	expression := &expressions.PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Value,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseInfixExpression(left expressions.Expression) expressions.Expression {
	expression := &expressions.InfixExpression{
		Token:    p.currentToken,
		Left:     left,
		Operator: p.currentToken.Value,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

var precedences = map[token.Type]int{
	token.Equal:              EQUALS,
	token.NotEqual:           EQUALS,
	token.LessThan:           LESS_GREATER,
	token.GreaterThan:        LESS_GREATER,
	token.LessThanOrEqual:    LESS_GREATER,
	token.GreaterThanOrEqual: LESS_GREATER,
	token.Plus:               SUM,
	token.Minus:              SUM,
	token.Multiply:           PRODUCT,
	token.Divide:             PRODUCT,
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}
