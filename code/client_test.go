package code

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchData(t *testing.T) {
	id := "42"
	expected := "Hello, World"
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/data/"+id, r.URL.String())
			w.Write([]byte(expected))
		},
	))
	defer ts.Close()

	actual, err := FetchData(ts.URL, id)

	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}