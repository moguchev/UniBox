/*
 * Copyright (C) 2020. Leonid Moguchev
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
			"URL":         r.URL.Path,
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(models.CtxKey(models.ReqIDKey)),
		}).Info(r.URL.Path)

		next.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"method":      r.Method,
			"URL":         r.URL.Path,
			"remote_addr": r.RemoteAddr,
			"work_time":   time.Since(start),
			"request_id":  r.Context().Value(models.CtxKey(models.ReqIDKey)),
		}).Info(r.URL.Path)
	})
}
