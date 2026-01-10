package cartrepository

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

type CreateCartStore interface {
	CreateCart(
		ctx context.Context,
		data *cartmodel.CartCreate,
	) error
}

type createCartRepo struct {
	store CreateCartStore
}

func NewCreateCartRepo(store CreateCartStore) *createCartRepo {
	return &createCartRepo{store: store}
}

func (repo *createCartRepo) MakeCart(
	ctx context.Context,
	userId int,
) error {
	data := &cartmodel.CartCreate{
		UserId: userId,
	}
	return repo.store.CreateCart(ctx, data)
}
