package assert

import (
	"reflect"
	"testing"
)

func Equal[T any](t *testing.T, a, b T) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("Assertion Failed: <a> is not equal to <b>\n\twhere\n\t\ta is -> %v,\n\t\tb is -> %v", a, b)
	}
}
