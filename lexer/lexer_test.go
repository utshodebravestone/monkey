package lexer

import (
	"testing"

	"monkey/assert"
	"monkey/text"
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
	l := New("({[%*+-^/:,=;]})<>")

	expectedTokens := []token.Token{
		token.New(token.OPEN_PAREN, "(", text.NewSpan(0, 1)),
		token.New(token.OPEN_BRACE, "{", text.NewSpan(1, 2)),
		token.New(token.OPEN_BRACKET, "[", text.NewSpan(2, 3)),
		token.New(token.MODULO, "%", text.NewSpan(3, 4)),
		token.New(token.ASTERISK, "*", text.NewSpan(4, 5)),
		token.New(token.PLUS, "+", text.NewSpan(5, 6)),
		token.New(token.MINUS, "-", text.NewSpan(6, 7)),
		token.New(token.HAT, "^", text.NewSpan(7, 8)),
		token.New(token.SLASH, "/", text.NewSpan(8, 9)),
		token.New(token.COLON, ":", text.NewSpan(9, 10)),
		token.New(token.COMMA, ",", text.NewSpan(10, 11)),
		token.New(token.EQUAL, "=", text.NewSpan(11, 12)),
		token.New(token.SEMICOLON, ";", text.NewSpan(12, 13)),
		token.New(token.CLOSE_BRACKET, "]", text.NewSpan(13, 14)),
		token.New(token.CLOSE_BRACE, "}", text.NewSpan(14, 15)),
		token.New(token.CLOSE_PAREN, ")", text.NewSpan(15, 16)),
		token.New(token.LESSER, "<", text.NewSpan(16, 17)),
		token.New(token.GREATER, ">", text.NewSpan(17, 18)),
	}

	for _, expected := range expectedTokens {
		actual := l.NextToken()
		assert.Equal(t, expected, actual)
	}
}

func TestNextTokenForDoubleCharToken(t *testing.T) {
	l := New("!===>=<=")

	expectedTokens := []token.Token{
		token.New(token.BANG_EQUAL, "!=", text.NewSpan(0, 2)),
		token.New(token.EQUAL_EQUAL, "==", text.NewSpan(2, 4)),
		token.New(token.GREATER_EQUAL, ">=", text.NewSpan(4, 6)),
		token.New(token.LESSER_EQUAL, "<=", text.NewSpan(6, 8)),
	}

	for _, expected := range expectedTokens {
		actual := l.NextToken()
		assert.Equal(t, expected, actual)
	}
}

func TestNextTokenWithSpace(t *testing.T) {
	l := New("* ^")

	expectedTokens := []token.Token{
		token.New(token.ASTERISK, "*", text.NewSpan(0, 1)),
		token.New(token.HAT, "^", text.NewSpan(2, 3)),
	}

	for _, expected := range expectedTokens {
		actual := l.NextToken()
		assert.Equal(t, expected, actual)
	}
}

func TestNextTokenForKeywords(t *testing.T) {
	l := New("true false let fun ret if else not and or duh")

	expectedTokens := []token.Token{
		token.New(token.TRUE, "true", text.NewSpan(0, 4)),
		token.New(token.FALSE, "false", text.NewSpan(5, 10)),
		token.New(token.LET, "let", text.NewSpan(11, 14)),
		token.New(token.FUN, "fun", text.NewSpan(15, 18)),
		token.New(token.RET, "ret", text.NewSpan(19, 22)),
		token.New(token.IF, "if", text.NewSpan(23, 25)),
		token.New(token.ELSE, "else", text.NewSpan(26, 30)),
		token.New(token.NOT, "not", text.NewSpan(31, 34)),
		token.New(token.AND, "and", text.NewSpan(35, 38)),
		token.New(token.OR, "or", text.NewSpan(39, 41)),
		token.New(token.IDENTIFIER, "duh", text.NewSpan(42, 45)),
	}

	for _, expected := range expectedTokens {
		actual := l.NextToken()
		assert.Equal(t, expected, actual)
	}
}

func TestNextTokenForLiterals(t *testing.T) {
	l := New("a99\"valid\"\"!valid")

	expectedTokens := []token.Token{
		token.New(token.IDENTIFIER, "a", text.NewSpan(0, 1)),
		token.New(token.NUMBER, "99", text.NewSpan(1, 3)),
		token.New(token.STRING, "\"valid\"", text.NewSpan(3, 10)),
		token.New(token.ILLEGAL, "\"!valid", text.NewSpan(10, 17)),
	}

	for _, expected := range expectedTokens {
		actual := l.NextToken()
		assert.Equal(t, expected, actual)
	}
}

func TestNextTokenWithIllegal(t *testing.T) {
	l := New("~")

	expected := token.New(token.ILLEGAL, "~", text.NewSpan(0, 1))

	actual := l.NextToken()

	assert.Equal(t, expected, actual)
}
