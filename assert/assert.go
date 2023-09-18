package assert

import (
	"reflect"
	"testing"
)

func Equal[T any](t *testing.T, expected, actual T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Assertion Failed: <expected> is not equal to <actual>\n\twhere\n\t\texpected is \t:\t %v,\n\t\tactual is \t:\t %v", expected, actual)
	}
}
