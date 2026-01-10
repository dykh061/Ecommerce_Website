package cartrepository

import "context"

type DeleteCartItemStorage interface {
	DeleteCartItem(ctx context.Context, id int) error
}

type deleteCartItemRepo struct {
	storage DeleteCartItemStorage
}

func NewDeleteCartItemRepo(storage DeleteCartItemStorage) *deleteCartItemRepo {
	return &deleteCartItemRepo{storage: storage}
}

func (repo *deleteCartItemRepo) DeleteCartItem(
	ctx context.Context,
	id int,
) error {
	return repo.storage.DeleteCartItem(ctx, id)
}
