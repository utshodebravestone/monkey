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

func TestNextTokenForSingleCharToken(t *testing.T) {
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

func TestNextTokenWithSpace(t *testing.T) {
	l := New("* ^")

	expectedTokens := []token.Token{
		token.New(token.ASTERISK, "*", diagnostics.NewSpan(0, 1)),
		token.New(token.HAT, "^", diagnostics.NewSpan(2, 3)),
	}

	for _, expected := range expectedTokens {
		actual := l.NextToken()
		assert.Equal(t, expected, actual)
	}
}

func TestNextTokenForKeywords(t *testing.T) {
	l := New("true false let fun ret if else duh")

	expectedTokens := []token.Token{
		token.New(token.TRUE, "true", diagnostics.NewSpan(0, 4)),
		token.New(token.FALSE, "false", diagnostics.NewSpan(5, 10)),
		token.New(token.LET, "let", diagnostics.NewSpan(11, 14)),
		token.New(token.FUN, "fun", diagnostics.NewSpan(15, 18)),
		token.New(token.RET, "ret", diagnostics.NewSpan(19, 22)),
		token.New(token.IF, "if", diagnostics.NewSpan(23, 25)),
		token.New(token.ELSE, "else", diagnostics.NewSpan(26, 30)),
		token.New(token.IDENTIFIER, "duh", diagnostics.NewSpan(31, 34)),
	}

	for _, expected := range expectedTokens {
		actual := l.NextToken()
		assert.Equal(t, expected, actual)
	}
}

func TestNextTokenForLiterals(t *testing.T) {
	l := New("a99\"valid\"\"!valid")

	expectedTokens := []token.Token{
		token.New(token.IDENTIFIER, "a", diagnostics.NewSpan(0, 1)),
		token.New(token.NUMBER, "99", diagnostics.NewSpan(1, 3)),
		token.New(token.STRING, "\"valid\"", diagnostics.NewSpan(3, 10)),
		token.New(token.ILLEGAL, "\"!valid", diagnostics.NewSpan(10, 17)),
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
