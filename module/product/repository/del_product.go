package productrepository

import "context"

type DeleteProductStorage interface {
	DeleteProduct(
		ctx context.Context,
		productId int,
	) error
}

type deleteProductRepo struct {
	storage DeleteProductStorage
}

func NewDeleteProductRepo(
	storage DeleteProductStorage,
) *deleteProductRepo {
	return &deleteProductRepo{storage: storage}
}

func (repo *deleteProductRepo) DeleteProduct(
	ctx context.Context,
	productId int,
) error {
	return repo.storage.DeleteProduct(
		ctx,
		productId,
	)
}
