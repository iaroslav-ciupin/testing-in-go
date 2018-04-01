package code // package code_test

import "testing"

func TestAnswer(t *testing.T) {
	expected := 42

	actual, err := answer() // get answer over the network

	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}
}
