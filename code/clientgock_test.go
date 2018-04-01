package code

import (
	"testing"
	"net/http"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchDataWithGock(t *testing.T) {
	url := "http://example.com"
	id := "42"
	gock.New(url).
		Get("/data/"+id).
		Reply(http.StatusOK).
		BodyString("OK")
	defer gock.OffAll()

	actual, err := FetchData(url, id)

	require.NoError(t, err)
	assert.Equal(t, "OK", actual)
	assert.False(t, gock.HasUnmatchedRequest())
	assert.True(t, gock.IsDone())
}