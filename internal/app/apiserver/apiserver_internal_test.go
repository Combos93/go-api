package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig(map[string]string{}))
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	s.handleHello().ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "Hello")
	assert.HTTPBodyContains(t, s.handleHello(), "GET", "/hello", nil, "Hell")
	assert.HTTPSuccess(t, s.handleHello(), "POST", "/hello", nil)
}
