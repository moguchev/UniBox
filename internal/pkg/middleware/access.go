package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/moguchev/UniBox/internal/app/models"
)

// AccessLogMiddleware - тайминги запросов
func (mw *Middleware) AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.WithFields(log.Fields{
			"method":      r.Method,
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(models.CtxKey("rID")),
		}).Info(r.URL.Path)

		next.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"method":      r.Method,
			"remote_addr": r.RemoteAddr,
			"work_time":   time.Since(start),
			"request_id":  r.Context().Value(models.CtxKey("rID")),
		}).Info(r.URL.Path)
	})
}
