package usecase

import (
	"context"
	"time"

	"github.com/moguchev/UniBox/internal/app/models"
	"github.com/moguchev/UniBox/internal/app/users"
	"github.com/moguchev/UniBox/internal/pkg/messages"
	"github.com/moguchev/UniBox/internal/pkg/validator"
)

type userUsecase struct {
	usersRepo      users.Repository
	contextTimeout time.Duration
}

// NewUsersUsecase will create new an userUsecase object representation of user.Usecase interface
func NewUsersUsecase(repo users.Repository, timeout time.Duration) users.Usecase {
	return &userUsecase{
		usersRepo:      repo,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, user models.NewUser) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	// rID := ctx.Value(models.CtxKey(models.ReqIDKey))
	usr := models.User{}

	target, err := validator.ValidateNewUser(user)
	if err != nil {
		err = models.Error{
			Type:     models.Invalid,
			Target:   target,
			Message:  err.Error(),
			Original: err,
		}
		return usr, err
	}

	usr, err = u.usersRepo.CreateUser(ctx, user)
	if err != nil {
		errorType := models.NoType
		var msg string

		e := err.(models.Error)
		switch e.Type {
		case models.AlreadyExists, models.Invalid:
			errorType = e.Type
			target = e.Target
			msg = messages.ErrorToMessage[e.Type]
			break
		default:
			errorType = models.Internal
			msg = messages.ErrorToMessage[errorType]
			break
		}

		err = models.Error{
			Type:     errorType,
			Target:   target,
			Message:  msg,
			Original: e,
		}
	}

	return usr, err
}
