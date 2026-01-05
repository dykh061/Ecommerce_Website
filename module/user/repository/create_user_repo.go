package srrepository

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

type CreateUserStorage interface {
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

type createUserRepo struct {
	storage CreateUserStorage
}

func NewCreateUserRepo(storage CreateUserStorage) *createUserRepo {
	return &createUserRepo{storage: storage}
}

func (repo *createUserRepo) CreateUser(
	ctx context.Context,
	data *usermodel.UserCreate,
) error {
	return repo.storage.Create(ctx, data)
}
