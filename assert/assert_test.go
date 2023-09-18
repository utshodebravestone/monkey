package assert_test

import (
	"testing"

	"monkey/assert"
)

func TestEqual(t *testing.T) {
	a := 5
	b := 5

	assert.Equal(t, a, b)
}
