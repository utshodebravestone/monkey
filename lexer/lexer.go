package lexer

import (
	"monkey/diagnostics"
	"monkey/token"
)

type Lexer struct {
	input string

	readPosition int // current char's index
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

func (l *Lexer) skipWhitespaces() {
	for isWhitespace(l.currentChar) {
		l.readNextChar()
	}
}

func (l *Lexer) readIdentifier() {
	for isLetter(l.currentChar) {
		l.readNextChar()
	}
}

func (l *Lexer) readNumber() {
	for isDigit(l.currentChar) {
		l.readNextChar()
	}
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) newToken(kind token.TokenKind, startPosition int, endPosition int) token.Token {
	return token.New(kind, l.input[startPosition:endPosition], diagnostics.NewSpan(startPosition, endPosition))
}

func (l *Lexer) NextToken() token.Token {
	var tok = token.Token{}

	l.skipWhitespaces()

	startPosition := l.readPosition

	switch l.currentChar {
	case '+':
		tok = l.newToken(token.PLUS, startPosition, l.peekPosition)
	case '-':
		tok = l.newToken(token.MINUS, startPosition, l.peekPosition)
	case '*':
		tok = l.newToken(token.ASTERISK, startPosition, l.peekPosition)
	case '/':
		tok = l.newToken(token.SLASH, startPosition, l.peekPosition)
	case '%':
		tok = l.newToken(token.MODULO, startPosition, l.peekPosition)
	case '^':
		tok = l.newToken(token.HAT, startPosition, l.peekPosition)
	case '!':
		tok = l.newToken(token.BANG, startPosition, l.peekPosition)
	case '=':
		tok = l.newToken(token.EQUAL, startPosition, l.peekPosition)

	case '(':
		tok = l.newToken(token.OPEN_PAREN, startPosition, l.peekPosition)
	case ')':
		tok = l.newToken(token.CLOSE_PAREN, startPosition, l.peekPosition)
	case '{':
		tok = l.newToken(token.OPEN_BRACE, startPosition, l.peekPosition)
	case '}':
		tok = l.newToken(token.CLOSE_BRACE, startPosition, l.peekPosition)
	case '[':
		tok = l.newToken(token.OPEN_BRACKET, startPosition, l.peekPosition)
	case ']':
		tok = l.newToken(token.CLOSE_BRACKET, startPosition, l.peekPosition)
	case ',':
		tok = l.newToken(token.COMMA, startPosition, l.peekPosition)
	case ';':
		tok = l.newToken(token.SEMICOLON, startPosition, l.peekPosition)

	// since this token doesn't live in the actual input, we have to provide a handmade lexeme
	case 0:
		tok = token.New(token.EOF, "0", diagnostics.NewSpan(startPosition, l.peekPosition))

	default:
		if isLetter(l.currentChar) {
			l.readIdentifier()
			lexeme := l.input[startPosition:l.readPosition]
			return l.newToken(token.LookupKeyword(lexeme), startPosition, l.readPosition)
		} else if isDigit(l.currentChar) {
			l.readNumber()
			return l.newToken(token.NUMBER, startPosition, l.readPosition)
		} else {
			tok = l.newToken(token.ILLEGAL, startPosition, l.peekPosition)
		}
	}

	l.readNextChar()

	return tok
}
