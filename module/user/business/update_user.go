package userbusiness

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
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
	_, err := biz.storage.FindDataWithCondition(ctx, condition)
	if err != nil {
		return err
	}
	if err := biz.storage.Update(ctx, condition, data); err != nil {
		return err
	}
	return nil
}
