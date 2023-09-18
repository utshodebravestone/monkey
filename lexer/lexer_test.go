package lexer

import (
	"testing"

	"monkey/assert"
	"monkey/diagnostics"
	"monkey/token"
)

func TestNew(t *testing.T) {
	expected := &Lexer{
		input:        "~",
		readPosition: 0,
		peekPosition: 1,
		currentChar:  '~',
	}

	actual := New("~")

	assert.Equal(t, expected, actual)
}

func TestReadNextChar(t *testing.T) {
	expected := &Lexer{
		input:        "~>",
		readPosition: 1,
		peekPosition: 2,
		currentChar:  '>',
	}

	actual := New("~>")
	actual.readNextChar()

	assert.Equal(t, expected, actual)
}

func TestGetCurrentLexeme(t *testing.T) {
	l := New("~>")

	// Iteration 1
	expected := "~"

	actual := l.getCurrentLexeme()

	assert.Equal(t, expected, actual)

	// Iteration 2
	expected = ">"

	l.readNextChar()
	actual = l.getCurrentLexeme()

	assert.Equal(t, expected, actual)
}

func TestGetCurrentSpan(t *testing.T) {
	l := New("~>")

	// Iteration 1
	expected := diagnostics.NewSpan(0, 1)

	actual := l.getCurrentSpan()

	assert.Equal(t, expected, actual)

	// Iteration 2
	expected = diagnostics.NewSpan(1, 2)

	l.readNextChar()
	actual = l.getCurrentSpan()

	assert.Equal(t, expected, actual)
}

func TestNextToken(t *testing.T) {
	l := New("({[%*+-^/!,=;]})")

	expectedTokens := []token.Token{
		token.New(token.OPEN_PAREN, "(", diagnostics.NewSpan(0, 1)),
		token.New(token.OPEN_BRACE, "{", diagnostics.NewSpan(1, 2)),
		token.New(token.OPEN_BRACKET, "[", diagnostics.NewSpan(2, 3)),
		token.New(token.MODULO, "%", diagnostics.NewSpan(3, 4)),
		token.New(token.ASTERISK, "*", diagnostics.NewSpan(4, 5)),
		token.New(token.PLUS, "+", diagnostics.NewSpan(5, 6)),
		token.New(token.MINUS, "-", diagnostics.NewSpan(6, 7)),
		token.New(token.HAT, "^", diagnostics.NewSpan(7, 8)),
		token.New(token.SLASH, "/", diagnostics.NewSpan(8, 9)),
		token.New(token.BANG, "!", diagnostics.NewSpan(9, 10)),
		token.New(token.COMMA, ",", diagnostics.NewSpan(10, 11)),
		token.New(token.EQUAL, "=", diagnostics.NewSpan(11, 12)),
		token.New(token.SEMICOLON, ";", diagnostics.NewSpan(12, 13)),
		token.New(token.CLOSE_BRACKET, "]", diagnostics.NewSpan(13, 14)),
		token.New(token.CLOSE_BRACE, "}", diagnostics.NewSpan(14, 15)),
		token.New(token.CLOSE_PAREN, ")", diagnostics.NewSpan(15, 16)),
	}

	for _, expected := range expectedTokens {
		actual := l.NextToken()
		assert.Equal(t, expected, actual)
	}
}

func TestNextTokenWithIllegal(t *testing.T) {
	l := New("~")

	expected := token.New(token.ILLEGAL, "~", diagnostics.NewSpan(0, 1))

	actual := l.NextToken()

	assert.Equal(t, expected, actual)
}
