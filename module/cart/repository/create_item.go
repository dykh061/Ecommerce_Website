package cartrepository

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

type CreateItemStorage interface {
	CreateItem(
		ctx context.Context,
		item *cartmodel.CartItemCreate,
	) error
}

type createItemRepo struct {
	storage CreateItemStorage
}

func NewCreateItemRepo(storage CreateItemStorage) *createItemRepo {
	return &createItemRepo{storage: storage}
}

func (repo *createItemRepo) AddItemToCart(
	ctx context.Context,
	item *cartmodel.CartItemCreate,
) error {
	return repo.storage.CreateItem(ctx, item)
}
