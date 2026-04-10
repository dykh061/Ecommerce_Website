package common

import (
	"errors"

	"gorm.io/gorm"
)

// IsRecordNotFound centralizes not-found detection so business code does not depend
// on ORM-specific sentinel errors.
func IsRecordNotFound(err error) bool {
	if err == nil {
		return false
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		return IsRecordNotFound(appErr.RootError())
	}

	return errors.Is(err, gorm.ErrRecordNotFound)
}
