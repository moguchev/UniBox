package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moguchev/UniBox/internal/pkg/middleware"
	"github.com/stretchr/testify/assert"
)

func TestCORSMiddleware(t *testing.T) {
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
