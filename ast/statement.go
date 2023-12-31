package ast

import (
	"fmt"

	"monkey/text"
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

func (ls *LetStatement) Span() text.Span {
	return text.ExtentSpan(ls.Token.Span, ls.Value.Span())
}

func (ls *LetStatement) ToString() string {
	if ls.Value != nil {
		return fmt.Sprintf("%s %s = %s;", ls.Token.Lexeme, ls.Name.ToString(), ls.Value.ToString())
	} else {
		return fmt.Sprintf("%s %s;", ls.Token.Lexeme, ls.Name.ToString())
	}
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) statement() {}

func (rs *ReturnStatement) Span() text.Span {
	return text.ExtentSpan(rs.Token.Span, rs.Value.Span())
}

func (rs *ReturnStatement) ToString() string {
	if rs.Value != nil {
		return fmt.Sprintf("%s %s;", rs.Token.Lexeme, rs.Value.ToString())
	} else {
		return fmt.Sprintf("%s;", rs.Token.Lexeme)
	}
}
