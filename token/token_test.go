package token_test

import (
	"testing"

	"monkey/assert"
	"monkey/diagnostics"
	"monkey/token"
)

func TestToString(t *testing.T) {
	expected := "Token `~` of Type ILLEGAL in 0:1"

	token := token.New(token.ILLEGAL, "~", diagnostics.NewSpan(0, 1))
	actual := token.ToString()

	assert.Equal(t, expected, actual)
}
