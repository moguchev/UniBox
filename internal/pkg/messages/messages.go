package messages

import "github.com/moguchev/UniBox/internal/app/models"

const (
	// Invalid - invalid
	Invalid = "invalid"
	// AlreadyExists - already exists
	AlreadyExists = "already exists"
	// NotFound - not found
	NotFound = "not found"
	// Internal -
	Internal = "internal"
)

// ErrorToMessage - mapping error to message
var ErrorToMessage map[models.ErrorType]string

func init() {
	ErrorToMessage = map[models.ErrorType]string{
		models.AlreadyExists: AlreadyExists,
		models.Invalid:       Invalid,
		models.Internal:      Internal,
	}
}
