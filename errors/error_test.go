package errors

import (
	"testing"

	"monkey/assert"
	"monkey/text"
)

func TestParserError(t *testing.T) {
	expected := "parse error: foo in 0:0"
	e := NewParseError("foo", text.NewSpan(0, 0))
	actual := e.Info()

	assert.Equal(t, expected, actual)
}
