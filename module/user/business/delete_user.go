package userbusiness

import (
	usermodel "OpenMarket/module/user/model"
	"context"
	"errors"
)

type DeleteUserStore interface {
	FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	Delete(ctx context.Context, id int) error
}

type deleteBusiness struct {
	store DeleteUserStore
}

func NewDeleteUserBusiness(store DeleteUserStore) *deleteBusiness {
	return &deleteBusiness{store: store}
}

func (biz *deleteBusiness) DeleteUser(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldData.Status == "deleted" {
		return errors.New("Data has been deleted")
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
