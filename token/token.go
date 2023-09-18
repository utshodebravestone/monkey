package token

import (
	"fmt"

	"monkey/diagnostics"
)

type Token struct {
	Kind   TokenKind
	Lexeme string
	Span   diagnostics.Span
}

func New(kind TokenKind, lexeme string, span diagnostics.Span) Token {
	return Token{
		Kind:   kind,
		Lexeme: lexeme,
		Span:   span,
	}
}

func (t *Token) ToString() string {
	return fmt.Sprintf("Token `%s` of Type %s in %s", t.Lexeme, t.Kind, t.Span.ToString())
}
