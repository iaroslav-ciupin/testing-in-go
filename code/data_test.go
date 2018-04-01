package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestF(t *testing.T) {
	for _, testCase := range []struct {
		input int
		expected int
	} {
		{input: 1, expected: 2},
		{input: -1, expected: 0},
		{input: 42, expected: 44},
	} {
		assert.Equal(t, testCase.expected, f(testCase.input), "input %d", testCase.input)
	}
}