package parser

import (
	"fmt"
	"monkey/assert"
	"monkey/errors"
	"monkey/lexer"
	"monkey/token"
	"testing"
)

func TestReadNextToken(t *testing.T) {
	input := `foo 1`

	l := lexer.New(input)
	p := New(l)

	assert.Equal(t, p.currentToken.Kind, token.IDENTIFIER)
	assert.Equal(t, p.currentToken.Lexeme, "foo")

	p.readNextToken()

	assert.Equal(t, p.currentToken.Kind, token.NUMBER)
	assert.Equal(t, p.currentToken.Lexeme, "1")
}

func TestAssertionMethods(t *testing.T) {
	input := `foo 1`

	l := lexer.New(input)
	p := New(l)

	assert.Equal(t, p.currentTokenIs(token.IDENTIFIER), true)
	assert.Equal(t, p.peekTokenIs(token.NUMBER), true)

	assert.Equal(t, p.expectPeek(token.NUMBER), true)

	assert.Equal(t, p.currentTokenIs(token.NUMBER), true)
	assert.Equal(t, p.peekTokenIs(token.EOF), true)
}

func TestParseLetStatement(t *testing.T) {
	input := `let foo = bar;`

	l := lexer.New(input)
	p := New(l)

	actual := p.parseLetStatement()

	assert.Equal(t, token.LET, actual.Token.Kind)
	assert.Equal(t, token.IDENTIFIER, actual.Name.Token.Kind)
	assert.Equal(t, "foo", actual.Name.Token.Lexeme)
}

func TestParseLetStatementFailure(t *testing.T) {
	input := `let foo bar;`

	l := lexer.New(input)
	p := New(l)

	actual := p.parseLetStatement()

	assert.Equal(t, nil, actual)

	expected := []errors.ParseError{
		errors.NewParseError(fmt.Sprintf("expected the token to be %s, but got %s instead", token.EQUAL, token.IDENTIFIER), p.peekToken.Span),
	}

	assert.Equal(t, p.Errors(), expected)
}

func TestParseProgram(t *testing.T) {
	input := `let foo = bar;`

	l := lexer.New(input)
	p := New(l)

	actual := p.ParseProgram()

	assert.Equal(t, len(actual.Statements), 1)
}
