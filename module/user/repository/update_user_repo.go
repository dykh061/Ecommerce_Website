package srrepository

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
)

type UpdateUserStorage interface {
	Update(ctx context.Context, condition map[string]interface{}, data usermodel.UserUpdate) error
}

type updateUserRepo struct {
	storage UpdateUserStorage
}

func NewUpdateUserRepo(storage UpdateUserStorage) *updateUserRepo {
	return &updateUserRepo{storage: storage}
}

func (repo *updateUserRepo) UpdateUserById(
	ctx context.Context,
	id int,
	data usermodel.UserUpdate,
) error {
	return repo.storage.Update(ctx, map[string]interface{}{
		"id":     id,
		"status": common.SystemStatusActive,
	}, data)
}
