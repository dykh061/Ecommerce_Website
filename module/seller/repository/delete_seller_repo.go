package sellerrepository

import "context"

type DeleteSellerStorage interface {
	DeleteSeller(ctx context.Context, id int) error
}

type deleteSellerRepo struct {
	storage DeleteSellerStorage
}

func NewDeleteSellerRepo(storage DeleteSellerStorage) *deleteSellerRepo {
	return &deleteSellerRepo{storage: storage}
}

func (repo *deleteSellerRepo) DeleteSellerByUserId(
	ctx context.Context,
	userId int,
) error {
	return repo.storage.DeleteSeller(ctx, userId)
}
