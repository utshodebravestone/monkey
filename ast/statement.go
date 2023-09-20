package ast

import (
	"fmt"

	"monkey/diagnostics"
	"monkey/token"
)

type Statement interface {
	Node
	statement()
}

type LetStatement struct {
	Token token.Token
	Name  *IdentifierExpression
	Value Expression
}

func (ls *LetStatement) statement() {}

func (ls *LetStatement) Span() diagnostics.Span {
	return diagnostics.ExtentSpan(ls.Token.Span, ls.Value.Span())
}

func (ls *LetStatement) ToString() string {
	return fmt.Sprintf("%s %s = %s", ls.Token.Lexeme, ls.Name.ToString(), ls.Value.ToString())
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) statement() {}

func (rs *ReturnStatement) Span() diagnostics.Span {
	return diagnostics.ExtentSpan(rs.Token.Span, rs.Value.Span())
}

func (rs *ReturnStatement) ToString() string {
	return fmt.Sprintf("%s %s", rs.Token.Lexeme, rs.Value.ToString())
}
