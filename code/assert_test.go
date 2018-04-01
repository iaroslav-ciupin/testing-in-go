package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssertAnswer(t *testing.T) {
	expect := assert.New(t)
	must := require.New(t)

	actual, err := answer()

	must.NoError(err)
	expect.Equal(42, actual)
}