package diagnostics

import (
	"testing"

	"monkey/assert"
)

func TestToString(t *testing.T) {
	expected := "0:1"

	span := NewSpan(0, 1)
	actual := span.ToString()

	assert.Equal(t, expected, actual)
}

func TestExtendSpan(t *testing.T) {
	expected := NewSpan(0, 2)

	span_a := NewSpan(0, 1)
	span_b := NewSpan(1, 2)
	actual := ExtentSpan(span_a, span_b)

	assert.Equal(t, expected, actual)
}
