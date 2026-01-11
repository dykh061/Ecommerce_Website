package srrepository

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

type UserStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)

	FindAddressById(
		ctx context.Context,
		id, userId int,
	) (*usermodel.UserAddress, error)
}
