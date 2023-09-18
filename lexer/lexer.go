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
	l.peekPosition++
}

func (l *Lexer) getCurrentLexeme() string {
	return l.input[l.readPosition:l.peekPosition]
}

func (l *Lexer) getCurrentSpan() diagnostics.Span {
	return diagnostics.NewSpan(uint(l.readPosition), uint(l.peekPosition))
}

func (l *Lexer) NextToken() token.Token {
	var tok = token.Token{}

	switch l.currentChar {
	default:
		tok = token.New(token.ILLEGAL, l.getCurrentLexeme(), l.getCurrentSpan())
	}

	return tok
}
