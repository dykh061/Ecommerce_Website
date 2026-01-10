package cartrepository

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

type FindCartStorage interface {
	FindCart(
		ctx context.Context,
		userId int,
	) (*cartmodel.Cart, error)
}

type findCartRepo struct {
	storage FindCartStorage
}

func NewFindCartRepo(storage FindCartStorage) *findCartRepo {
	return &findCartRepo{storage: storage}
}

func (repo *findCartRepo) FindCartWithId(
	ctx context.Context,
	userId int,
) (*cartmodel.Cart, error) {
	return repo.storage.FindCart(ctx, userId)
}
