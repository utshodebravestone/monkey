package ast

import (
	"monkey/text"
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

func (ie *IdentifierExpression) Span() text.Span {
	return ie.Token.Span
}

func (ie *IdentifierExpression) ToString() string {
	return ie.Token.Lexeme
}
