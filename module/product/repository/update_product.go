package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type UpdateProductStorage interface {
	UpdateProduct(
		ctx context.Context,
		condition map[string]interface{},
		data *productmodel.ProductUpdate,
	) error
}

type updateProductRepo struct {
	storage UpdateProductStorage
}

func NewUpdateProductRepo(storage UpdateProductStorage) *updateProductRepo {
	return &updateProductRepo{storage: storage}
}

func (repo *updateProductRepo) UpdateProduct(
	ctx context.Context,
	productID int,
	condition map[string]interface{},
	data *productmodel.ProductUpdate,
) error {
	// Update product fields
	if err := repo.storage.UpdateProduct(ctx, condition, data); err != nil {
		return err
	}

	return nil
}
