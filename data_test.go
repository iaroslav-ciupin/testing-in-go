package main

import "testing"

func TestF(t *testing.T) {
	for _, testCase := range []struct {
		input int
		expected int
	} {
		{
			input: 1, 
			expected: 2,
		},
		{
			input: 2, 
			expected: 3,
		},
	} {
		if testCase.expected != f(testCase.input) {
			t.Errorf("test failed %d", testCase.input)
		}
	}
}