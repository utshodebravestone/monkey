package token

import (
	"fmt"

	"monkey/text"
)

type Token struct {
	Kind   TokenKind
	Lexeme string
	Span   text.Span
}

func New(kind TokenKind, lexeme string, span text.Span) Token {
	return Token{
		Kind:   kind,
		Lexeme: lexeme,
		Span:   span,
	}
}

func (t *Token) ToString() string {
	return fmt.Sprintf("Token `%s` of Type %s in %s", t.Lexeme, t.Kind, t.Span.ToString())
}
