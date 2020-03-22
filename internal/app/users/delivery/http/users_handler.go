package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/moguchev/UniBox/internal/app/users"
)

// UsersHandler represent the httphandler for users
type UsersHandler struct {
	UsersUsecase users.Usecase
}

// NewUsersHandler will initialize the user(s)/ resources endpoint
func NewUsersHandler(router *mux.Router, us users.Usecase) {
	handler := &UsersHandler{
		UsersUsecase: us,
	}

	router.HandleFunc("/user", handler.CreateUser).Methods(http.MethodPost, http.MethodOptions)
}

// CreateUser - handler
func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("[1,2,3]"))
}
