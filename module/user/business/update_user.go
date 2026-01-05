package userbusiness

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
	"errors"
)

type UpdateUserRepo interface {
	UpdateUserById(
		ctx context.Context,
		id int,
		data usermodel.UserUpdate,
	) error
}

type updateUserBusiness struct {
	repo   UpdateUserRepo
	finder ActiveUserFinder
}

func NewUpdateUserBusiness(
	repo UpdateUserRepo,
	finder ActiveUserFinder,
) *updateUserBusiness {
	return &updateUserBusiness{
		repo:   repo,
		finder: finder,
	}
}

func (biz *updateUserBusiness) UpdateUser(ctx context.Context, id int, data usermodel.UserUpdate) error {
	oldData, err := biz.finder.FindActiveUserByID(ctx, id)
	if err != nil {
		return common.ErrorDB(err)
	}
	if oldData == nil {
		return common.ErrEntityNotFound(usermodel.EntityName, errors.New("user not found"))
	}
	if oldData.IsBanned {
		return common.ErrInvalidState(usermodel.EntityName, "banned")
	}
	if err := biz.repo.UpdateUserById(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}
