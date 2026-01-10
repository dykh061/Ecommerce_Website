package cartrepository

import (
	"context"
)

type UpdateCartItemStore interface {
	UpdateCartItem(
		ctx context.Context,
		id, by int,
	) error
}

type updateCartItemRepo struct {
	store UpdateCartItemStore
}

func NewUpdateCartItemRepo(store UpdateCartItemStore) *updateCartItemRepo {
	return &updateCartItemRepo{store: store}
}

func (repo *updateCartItemRepo) AdjustCartItem(
	ctx context.Context,
	id, quantity int,
) error {
	return repo.store.UpdateCartItem(ctx, id, quantity)
}
