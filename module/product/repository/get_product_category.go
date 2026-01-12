package productrepository

import "context"

type GetProductCategoryStorage interface {
	GetProductCategoryID(
		ctx context.Context,
		productID int,
	) (*int, error)
}

type getProductCategoryRepo struct {
	storage GetProductCategoryStorage
}

func NewGetProductCategoryRepo(storage GetProductCategoryStorage) *getProductCategoryRepo {
	return &getProductCategoryRepo{storage: storage}
}

func (repo *getProductCategoryRepo) GetProductCategoryID(
	ctx context.Context,
	productID int,
) (*int, error) {
	return repo.storage.GetProductCategoryID(ctx, productID)
}
