package parser

import (
	"monkey/assert"
	"monkey/lexer"
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `foo 1`

	l := lexer.New(input)
	p := New(l)

	assert.Equal(t, p.currentToken.Kind, token.IDENTIFIER)
	assert.Equal(t, p.currentToken.Lexeme, "foo")

	p.nextToken()

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
