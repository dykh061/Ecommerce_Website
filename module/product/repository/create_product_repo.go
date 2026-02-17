package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type CreateProductStorage interface {
	Create(ctx context.Context, data *productmodel.ProductCreate) error
}

type createProductRepo struct {
	productstore CreateProductStorage
}

func NewCreateProductRepo(pstore CreateProductStorage) *createProductRepo {
	return &createProductRepo{productstore: pstore}
}

func (repo *createProductRepo) CreateProduct(
	ctx context.Context,
	data *productmodel.ProductCreate,
) (*productmodel.ProductCreate, error) {
	if err := repo.productstore.Create(ctx, data); err != nil {
		return nil, err
	}
	data.Mask()
	return data, nil
}
