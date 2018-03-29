package main

import "testing"

func f() int {
	return 42
}

func TestValidate(t *testing.T) {
	expected := 42

	actual := f()

	if expected != actual {
		t.Error("expected %d, but got %d", expected, actual)
	}
}
