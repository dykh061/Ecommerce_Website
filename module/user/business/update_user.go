package userbusiness

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

type UpdateUserStorage interface {
	Update(ctx context.Context, id int, data usermodel.UserUpdate) error
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
	_, err := biz.storage.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if err := biz.storage.Update(ctx, id, data); err != nil {
		return err
	}
	return nil
}
