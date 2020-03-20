package handler

import "net/http"

type Handler struct {
	InternalDir string
	// AuthService session.ServiceClient
	// UserService users.Service
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("[1,2,3]"))
}
