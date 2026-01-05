package srrepository

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

type ChangePasswordStorage interface {
	Update(ctx context.Context, condition map[string]interface{}, data usermodel.UserUpdate) error
}

type changePasswordRepo struct {
	store ChangePasswordStorage
}

func NewChangePasswordRepo(storage ChangePasswordStorage) *changePasswordRepo {
	return &changePasswordRepo{store: storage}

}

func (repo *changePasswordRepo) UpdatePassword(
	ctx context.Context,
	userID int,
	hashedPassword string,
) error {

	update := usermodel.UserUpdate{
		Password: &hashedPassword,
	}

	return repo.store.Update(ctx, map[string]interface{}{
		"id": userID,
	}, update)
}
