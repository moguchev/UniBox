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

package respond

import (
	"encoding/json"
	"net/http"

	"github.com/moguchev/UniBox/internal/app/models"
)

// Error - answer with error log
func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	if e, ok := err.(models.Error); ok {
		msg := models.ErrorMessage{
			Target:      e.Target,
			Message:     e.Message,
			ContextInfo: e.ContextInfo,
		}
		Respond(w, r, code, msg)
	} else {
		Respond(w, r, code, err)
	}
}

// Respond - http json respond
func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
}
