package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/moguchev/UniBox/internal/app/models"
	"github.com/moguchev/UniBox/internal/app/users"
	respond "github.com/moguchev/UniBox/pkg/respond"
)

// UsersHandler represent the httphandler for users
type UsersHandler struct {
	Usecase users.Usecase
}

// NewUsersHandler will initialize the user(s)/ resources endpoint
func NewUsersHandler(router *mux.Router, us users.Usecase) {
	handler := &UsersHandler{
		Usecase: us,
	}

	router.HandleFunc("/user", handler.CreateUser).Methods(
		http.MethodPost,
		http.MethodOptions)
}

// CreateUser - handler
func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	rID := ctx.Value(models.CtxKey("rID"))

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = models.Error{
			Type:     models.BadRequest,
			Target:   "body",
			Message:  "invalid",
			Original: err,
		}
		respond.Error(w, r, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		err = models.Error{
			Type:     models.BadRequest,
			Target:   "body",
			Message:  "invalid",
			Original: err,
		}
		respond.Error(w, r, http.StatusBadRequest, err)
		return
	}

	log.WithFields(log.Fields{
		"request_id": rID,
		"user":       user,
	}).Debug("Unmarshal user")

	err = h.Usecase.CreateUser(ctx, user)
	if err != nil {
		respond.Error(w, r, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(bytes)
}
