package userbusiness

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
	"errors"
)

type UpdateUserStorage interface {
	Update(ctx context.Context, condition map[string]interface{}, data usermodel.UserUpdate) error
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type updateUserBusiness struct {
	storage UpdateUserStorage
}

func NewUpdateUserBusiness(storage UpdateUserStorage) *updateUserBusiness {
	return &updateUserBusiness{storage: storage}
}

func (biz *updateUserBusiness) UpdateUser(ctx context.Context, id int, data usermodel.UserUpdate) error {
	condition := map[string]interface{}{
		"id":     id,
		"status": common.SystemStatusActive,
	}
	oldData, err := biz.storage.FindDataWithCondition(ctx, condition)
	if err != nil {
		return common.ErrorDB(err)
	}
	if oldData == nil {
		return common.ErrEntityNotFound(usermodel.EntityName, errors.New("user not found"))
	}
	if oldData.IsBanned {
		return common.ErrInvalidState(usermodel.EntityName, "banned")
	}

	if err := biz.storage.Update(ctx, condition, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}
