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
