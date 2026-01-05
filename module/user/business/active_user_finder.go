package userbusiness

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

type ActiveUserFinder interface {
	FindActiveUserByID(
		ctx context.Context,
		userID int,
	) (*usermodel.User, error)
}
type AciveUEmailFinder interface {
	FindUserWithEmail(
		ctx context.Context,
		email string,
	) (*usermodel.User, error)
}
