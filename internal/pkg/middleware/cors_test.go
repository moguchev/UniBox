package middleware_test

import (
	"net/http"
	"net/http/httptest"
	s "strings"
	"testing"

	"github.com/moguchev/UniBox/internal/pkg/middleware"
	"github.com/stretchr/testify/assert"
)

func TestCORSMiddlewareOrigin(t *testing.T) {
	mw := middleware.InitMiddleware()
	handler := mw.CORSMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	assert.Equal(t, "*", res.Header().Get("Access-Control-Allow-Origin"))
}

func TestCORSMiddlewareOptions(t *testing.T) {
	mw := middleware.InitMiddleware()
	handler := mw.CORSMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	assert.Equal(t, "*", res.Header().Get("Access-Control-Allow-Origin"))

	methods := res.Header().Get("Access-Control-Allow-Methods")
	assert.NotEmpty(t, methods)
	assert.True(t, s.Contains(methods, http.MethodGet))
	assert.True(t, s.Contains(methods, http.MethodPost))
	assert.True(t, s.Contains(methods, http.MethodPut))
	assert.True(t, s.Contains(methods, http.MethodDelete))
	assert.True(t, s.Contains(methods, http.MethodHead))
	assert.True(t, s.Contains(methods, http.MethodHead))

	headers := res.Header().Get("Access-Control-Allow-Headers")
	assert.NotEmpty(t, headers)
	assert.True(t, s.Contains(headers, "Content-Type"))
	assert.True(t, s.Contains(headers, "X-Content-Type-Options"))
	assert.True(t, s.Contains(headers, "X-Csrf-Token"))

	credentials := res.Header().Get("Access-Control-Allow-Credentials")
	assert.NotEmpty(t, credentials)
	assert.Equal(t, credentials, "true")
}
