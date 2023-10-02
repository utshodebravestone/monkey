package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/errors"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token

	errors []errors.ParseError
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.readNextToken() // loading currentToken
	p.readNextToken() // loading peekToken

	return p
}

func (p *Parser) readNextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) currentTokenIs(kind token.TokenKind) bool {
	return p.currentToken.Kind == kind
}

func (p *Parser) peekTokenIs(kind token.TokenKind) bool {
	return p.peekToken.Kind == kind
}

func (p *Parser) expectPeek(kind token.TokenKind) bool {
	if p.peekTokenIs(kind) {
		p.readNextToken()
		return true
	} else {
		p.errors = append(p.errors, errors.NewParseError(fmt.Sprintf("expected the token to be %s, but got %s instead", kind, p.peekToken.Kind), p.peekToken.Span))
		return false
	}
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Kind {
	case token.LET:
		return p.parseLetStatement()
	case token.RET:
		return p.parseRetStatement()

	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.IdentifierExpression{
		Token: p.currentToken,
	}

	if !p.expectPeek(token.EQUAL) {
		return nil
	}

	for !p.currentTokenIs(token.SEMICOLON) {
		p.readNextToken()
	}

	return stmt
}

func (p *Parser) parseRetStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	for !p.currentTokenIs(token.SEMICOLON) {
		p.readNextToken()
	}

	return stmt
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.currentToken.Kind != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.readNextToken()
	}

	return program
}

func (p *Parser) Errors() []errors.ParseError {
	return p.errors
}
