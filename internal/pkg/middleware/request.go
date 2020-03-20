package middleware

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
)

type ctxKey string

// RequestIDMiddleware - присвоение запросу id
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := strconv.Itoa(rand.Int())
		ctx := r.Context()
		ctx = context.WithValue(ctx, ctxKey("rID"), reqID)
		w.Header().Set("Request-ID", reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
