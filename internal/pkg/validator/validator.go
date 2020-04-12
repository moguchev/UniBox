package validator

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/moguchev/UniBox/internal/app/models"
	"github.com/moguchev/UniBox/internal/pkg/messages"
)

const (
	nickRgx = "^[A-Za-z0-9!#$%&'*+=?^_`{|}~.-]{3,20}$"
	nameRgx = "^[A-Za-zА-Яа-яЁё-]{1,50}$"
)

// ValidateNewUser -
func ValidateNewUser(u models.NewUser) (string, error) {
	if ok := ValidateEmail(u.Email); !ok {
		return "email", errors.New(messages.Invalid)
	}
	if ok := ValidateNick(u.Nickname); !ok {
		return "nickname", errors.New(messages.Invalid)
	}
	if ok := ValidateFirstname(u.Firstname); !ok {
		return "firstname", errors.New(messages.Invalid)
	}
	if ok := ValidateLastname(u.Lastname); !ok {
		return "lastname", errors.New(messages.Invalid)
	}
	if ok := ValidatePassword(u.Password); !ok {
		return "password", errors.New(messages.Invalid)
	}
	return "", nil
}

// IsEmpty -
func IsEmpty(str string) bool {
	return len(str) == 0
}

// ValidateEmail -
func ValidateEmail(email string) bool {
	return govalidator.IsEmail(email)
}

// ValidateNick -
func ValidateNick(nick string) bool {
	return govalidator.Matches(nick, nickRgx)
}

// ValidateFirstname -
func ValidateFirstname(str string) bool {
	return govalidator.Matches(str, nameRgx)
}

// ValidateLastname -
func ValidateLastname(str string) bool {
	return govalidator.Matches(str, nameRgx)
}

// ValidatePassword -
func ValidatePassword(pass string) bool {
	return !IsEmpty(pass)
}
