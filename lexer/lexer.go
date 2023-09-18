package lexer

import (
	"monkey/diagnostics"
	"monkey/token"
)

type Lexer struct {
	input string

	readPosition int // currentChar's index
	peekPosition int // what we're going to read next

	currentChar byte // what we've just read
}

func New(input string) *Lexer {
	l := &Lexer{
		input:        input,
		readPosition: 0,
		peekPosition: 0,
		currentChar:  0,
	}

	l.readNextChar()

	return l
}

func (l *Lexer) readNextChar() {
	if l.peekPosition < len(l.input) {
		l.currentChar = l.input[l.peekPosition]
	} else {
		l.currentChar = 0
	}

	l.readPosition = l.peekPosition
	l.peekPosition += 1
}

func (l *Lexer) getCurrentLexeme() string {
	return l.input[l.readPosition:l.peekPosition]
}

func (l *Lexer) getCurrentSpan() diagnostics.Span {
	return diagnostics.NewSpan(uint(l.readPosition), uint(l.peekPosition))
}

func (l *Lexer) skipWhitespaces() {
	for isWhitespace(l.currentChar) {
		l.readNextChar()
	}
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func (l *Lexer) newToken(kind token.TokenKind) token.Token {
	return token.New(kind, l.getCurrentLexeme(), l.getCurrentSpan())
}

func (l *Lexer) NextToken() token.Token {
	var tok = token.Token{}

	l.skipWhitespaces()

	switch l.currentChar {
	case '+':
		tok = l.newToken(token.PLUS)
	case '-':
		tok = l.newToken(token.MINUS)
	case '*':
		tok = l.newToken(token.ASTERISK)
	case '/':
		tok = l.newToken(token.SLASH)
	case '%':
		tok = l.newToken(token.MODULO)
	case '^':
		tok = l.newToken(token.HAT)
	case '!':
		tok = l.newToken(token.BANG)
	case '=':
		tok = l.newToken(token.EQUAL)

	case '(':
		tok = l.newToken(token.OPEN_PAREN)
	case ')':
		tok = l.newToken(token.CLOSE_PAREN)
	case '{':
		tok = l.newToken(token.OPEN_BRACE)
	case '}':
		tok = l.newToken(token.CLOSE_BRACE)
	case '[':
		tok = l.newToken(token.OPEN_BRACKET)
	case ']':
		tok = l.newToken(token.CLOSE_BRACKET)
	case ',':
		tok = l.newToken(token.COMMA)
	case ';':
		tok = l.newToken(token.SEMICOLON)

	// since this token doesn't live in the actual input, we have to provide a handmade lexeme
	case 0:
		tok = token.New(token.EOF, "0", l.getCurrentSpan())

	default:
		tok = l.newToken(token.ILLEGAL)
	}

	l.readNextChar()

	return tok
}
