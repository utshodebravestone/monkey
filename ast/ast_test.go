package ast

import (
	"testing"

	"monkey/assert"
	"monkey/text"
	"monkey/token"
)

func TestIdentifierExpressionToString(t *testing.T) {
	ie := &IdentifierExpression{
		Token: token.Token{
			Kind:   token.IDENTIFIER,
			Lexeme: "foo",
			Span:   text.Span{}, // don't care about span since it was tested on lexer
		},
	}

	expected := "foo"

	actual := ie.ToString()

	assert.Equal(t, expected, actual)
}

func TestLetStatementToString(t *testing.T) {
	ls := LetStatement{
		Token: token.Token{
			Kind:   token.LET,
			Lexeme: "let",
			Span:   text.Span{}, // don't care about span since it was tested on lexer
		},
		Name: &IdentifierExpression{
			Token: token.Token{
				Kind:   token.IDENTIFIER,
				Lexeme: "foo",
				Span:   text.Span{}, // don't care about span since it was tested on lexer
			},
		},
		Value: &IdentifierExpression{
			Token: token.Token{
				Kind:   token.IDENTIFIER,
				Lexeme: "bar",
				Span:   text.Span{}, // don't care about span since it was tested on lexer
			},
		},
	}

	expected := "let foo = bar"

	actual := ls.ToString()

	assert.Equal(t, expected, actual)
}

func TestReturnStatementToString(t *testing.T) {
	ls := ReturnStatement{
		Token: token.Token{
			Kind:   token.RET,
			Lexeme: "ret",
			Span:   text.Span{}, // don't care about span since it was tested on lexer
		},
		Value: &IdentifierExpression{
			Token: token.Token{
				Kind:   token.IDENTIFIER,
				Lexeme: "foo",
				Span:   text.Span{}, // don't care about span since it was tested on lexer
			},
		},
	}

	expected := "ret foo"

	actual := ls.ToString()

	assert.Equal(t, expected, actual)
}
