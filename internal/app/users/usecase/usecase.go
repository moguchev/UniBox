package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	log "github.com/sirupsen/logrus"

	"github.com/moguchev/UniBox/internal/app/models"
	"github.com/moguchev/UniBox/internal/app/users"
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
	rID := ctx.Value(models.CtxKey("rID"))

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		log.WithFields(log.Fields{
			"request_id": rID,
		}).Debug(err)
		return err
	}
	if govalidator.IsNull(user.Password) {
		err = errors.New("password: empty")
		log.WithFields(log.Fields{
			"request_id": rID,
		}).Debug(err)
		return err
	}

	err = u.usersRepo.CreateUser(ctx, user)
	if err != nil {

	}

	return err
}
