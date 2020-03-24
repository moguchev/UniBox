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
			e.Target,
			e.Message,
			e.ContextInfo,
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