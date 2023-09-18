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
	l := New("~")

	expected := token.New(token.ILLEGAL, "~", diagnostics.NewSpan(0, 1))

	actual := l.NextToken()

	assert.Equal(t, expected, actual)
}
