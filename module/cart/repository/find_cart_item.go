package cartrepository

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

type FindCartItemStorage interface {
	FindCartItem(
		ctx context.Context,
		cartId, variantId int,
	) (*cartmodel.CartItem, error)
}

type findCartItemRepo struct {
	storage FindCartItemStorage
}

func NewFindCartItemRepo(storage FindCartItemStorage) *findCartItemRepo {
	return &findCartItemRepo{storage: storage}
}

func (repo *findCartItemRepo) FindCartItemWithId(
	ctx context.Context,
	cartId, variantId int,
) (*cartmodel.CartItem, error) {
	return repo.storage.FindCartItem(ctx, cartId, variantId)
}
