package assert

import (
	"testing"
)

func TestEqual(t *testing.T) {
	a := 5
	b := 5

	Equal(t, a, b)
}
