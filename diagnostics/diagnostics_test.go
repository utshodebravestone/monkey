package diagnostics_test

import (
	"testing"

	"monkey/assert"
	"monkey/diagnostics"
)

func TestToString(t *testing.T) {
	expected := "0:1"

	span := diagnostics.NewSpan(0, 1)
	actual := span.ToString()

	assert.Equal(t, expected, actual)
}

func TestExtendSpan(t *testing.T) {
	expected := diagnostics.NewSpan(0, 2)

	span_a := diagnostics.NewSpan(0, 1)
	span_b := diagnostics.NewSpan(1, 2)
	actual := diagnostics.ExtentSpan(span_a, span_b)

	assert.Equal(t, expected, actual)
}
