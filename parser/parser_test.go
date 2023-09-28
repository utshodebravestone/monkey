package parser

import (
	"monkey/assert"
	"monkey/lexer"
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `foo bar`

	l := lexer.New(input)
	p := New(l)

	assert.Equal(t, p.currentToken.Kind, token.IDENTIFIER)
	assert.Equal(t, p.currentToken.Lexeme, "foo")

	p.nextToken()

	assert.Equal(t, p.currentToken.Kind, token.IDENTIFIER)
	assert.Equal(t, p.currentToken.Lexeme, "bar")
}
