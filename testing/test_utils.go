package testing

import (
	"reflect"
	"testing"
)

func AssertEquals(t *testing.T, expected any, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected value %v\n\nActual length %v", expected, actual)
	}
}
