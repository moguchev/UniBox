package middleware

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/moguchev/UniBox/internal/app/models"
)

// RequestIDMiddleware - присвоение запросу id
func (mw *Middleware) RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := strconv.Itoa(rand.Int())
		ctx := r.Context()
		ctx = context.WithValue(ctx, models.CtxKey("rID"), reqID)
		w.Header().Set("Request-ID", reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
