package userbusiness

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
	"errors"
)

type DeleteUserRepo interface {
	DeleteUserByID(
		ctx context.Context,
		id int,
	) error
}

type deleteBusiness struct {
	fstorage ActiveUserFinder
	dstorage DeleteUserRepo
}

func NewDeleteUserBusiness(
	fstorage ActiveUserFinder,
	dstorage DeleteUserRepo,
) *deleteBusiness {
	return &deleteBusiness{
		fstorage: fstorage,
		dstorage: dstorage,
	}
}

func (biz *deleteBusiness) DeleteUser(ctx context.Context, id int) error {
	oldData, err := biz.fstorage.FindActiveUserByID(ctx, id)
	if err != nil {
		return common.ErrorDB(err)
	}
	if oldData == nil {
		return common.ErrEntityNotFound(usermodel.EntityName, errors.New("user not found"))
	}
	if oldData.Status == common.SystemStatusDeleted {
		return common.ErrInvalidState(usermodel.EntityName, "deleted")
	}

	if err := biz.dstorage.DeleteUserByID(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(usermodel.EntityName, err)
	}
	return nil
}
