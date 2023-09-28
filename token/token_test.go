package token

import (
	"testing"

	"monkey/assert"
	"monkey/text"
)

func TestToString(t *testing.T) {
	expected := "Token `~` of Type ILLEGAL in 0:1"

	tok := New(ILLEGAL, "~", text.NewSpan(0, 1))
	actual := tok.ToString()

	assert.Equal(t, expected, actual)
}

func TestLookupKeyword(t *testing.T) {
	expected := []TokenKind{
		TRUE,
		FALSE,
		LET,
		FUN,
		RET,
		IF,
		ELSE,
		IDENTIFIER,
	}

	actual := []TokenKind{
		LookupKeyword("true"),
		LookupKeyword("false"),
		LookupKeyword("let"),
		LookupKeyword("fun"),
		LookupKeyword("ret"),
		LookupKeyword("if"),
		LookupKeyword("else"),
		LookupKeyword("whatever"),
	}

	assert.Equal(t, expected, actual)
}
