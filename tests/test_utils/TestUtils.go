package test_utils

import (
	"testing"
)

func AssertEqual(t *testing.T, expected any, actual any) {
	if expected != actual {
		t.Fatalf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
