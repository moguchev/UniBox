package handler

import "net/http"

// CreateUser - handler
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("[1,2,3]"))
}
