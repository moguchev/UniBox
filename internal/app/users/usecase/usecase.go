package usecase

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

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

func (u *userUsecase) CreateUser(ctx context.Context, user models.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	rID := ctx.Value(models.CtxKey(models.ReqIDKey))

	target, err := validator.ValidateCreateUser(user)
	if err != nil {
		err = models.Error{
			Type:     models.BadRequest,
			Target:   target,
			Message:  err.Error(),
			Original: err,
		}
		log.WithFields(log.Fields{
			"request_id": rID,
			"place":      "usecase",
			"action":     "validation",
		}).Debug(err)
		return err
	}

	err = u.usersRepo.CreateUser(ctx, user)
	if err != nil {
		if e, ok := err.(models.Error); ok {
			switch e.Type {
			case models.AlreadyExists:
				err = models.Error{
					Type:     models.AlreadyExists,
					Target:   e.Target,
					Message:  messages.AlreadyExists,
					Original: e,
				}
				break
			default:
				err = models.Error{
					Type:     models.Internal,
					Message:  e.Error(),
					Original: e,
				}
				break
			}
		}
	}

	return err
}
