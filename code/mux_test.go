package code

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestMux(t *testing.T) {
	r := NewRouter()
	request := httptest.NewRequest(http.MethodGet, "/data", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "OK", response.Body.String())
}