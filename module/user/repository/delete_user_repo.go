package srrepository

import "context"

type DeleteUserStorage interface {
	Delete(ctx context.Context, id int) error
}

type deleteUserRepo struct {
	storage DeleteUserStorage
}

func NewDeleteUserRepo(storage DeleteUserStorage) *deleteUserRepo {
	return &deleteUserRepo{storage: storage}
}

func (repo *deleteUserRepo) DeleteUserByID(
	ctx context.Context,
	id int,
) error {
	return repo.storage.Delete(ctx, id)
}
