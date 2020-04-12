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
	"github.com/moguchev/UniBox/internal/pkg/messages"
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
	rID := ctx.Value(models.CtxKey(models.ReqIDKey))

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = models.Error{
			Type:     models.BadRequest,
			Target:   "body",
			Message:  messages.Invalid,
			Original: err,
		}
		respond.Error(w, r, http.StatusBadRequest, err)
		return
	}

	newUser := models.NewUser{}
	err = json.Unmarshal(bytes, &newUser)
	if err != nil {
		err = models.Error{
			Type:     models.BadRequest,
			Target:   "body",
			Message:  messages.Invalid,
			Original: err,
		}
		respond.Error(w, r, http.StatusBadRequest, err)
		return
	}

	log.WithFields(log.Fields{
		"request_id": rID,
		"place":      "delivery",
		"action":     "unmarshal",
	}).Debug(newUser)

	err = h.Usecase.CreateUser(ctx, newUser)
	var code int
	if err != nil {
		e := err.(models.Error)
		switch e.Type {
		case models.AlreadyExists:
			code = http.StatusConflict
			break
		case models.Internal:
			code = http.StatusInternalServerError
			break
		default:
			code = http.StatusBadRequest
			break
		}
		respond.Error(w, r, code, e)
		return
	}

	user := models.User{
		Nickname:  newUser.Nickname,
		Email:     newUser.Email,
		Firstname: newUser.Firstname,
		Lastname:  newUser.Lastname,
	}
	bytes, _ = json.Marshal(user)
	w.WriteHeader(http.StatusCreated)
	w.Write(bytes)
}
