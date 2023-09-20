package ast

import (
	"monkey/diagnostics"
	"monkey/token"
)

type Expression interface {
	Node
	expression()
}

type IdentifierExpression struct {
	Token token.Token
}

func (ie *IdentifierExpression) expression() {}

func (ie *IdentifierExpression) Span() diagnostics.Span {
	return ie.Token.Span
}

func (ie *IdentifierExpression) ToString() string {
	return ie.Token.Lexeme
}
