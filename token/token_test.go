package token

import (
	"testing"

	"monkey/assert"
	"monkey/diagnostics"
)

func TestToString(t *testing.T) {
	expected := "Token `~` of Type ILLEGAL in 0:1"

	tok := New(ILLEGAL, "~", diagnostics.NewSpan(0, 1))
	actual := tok.ToString()

	assert.Equal(t, expected, actual)
}
