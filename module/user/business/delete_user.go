package userbusiness

import (
	"OpenMarket/common"
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
		return common.ErrorDB(err)
	}
	if oldData == nil {
		return common.ErrEntityNotFound(usermodel.EntityName, errors.New("user not found"))
	}
	if oldData.Status == common.SystemStatusDeleted {
		return common.ErrInvalidState(usermodel.EntityName, "deleted")
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(usermodel.EntityName, err)
	}
	return nil
}
