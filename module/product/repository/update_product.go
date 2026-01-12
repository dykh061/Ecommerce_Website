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
	UpsertProductCategory(
		ctx context.Context,
		productID int,
		categoryID int,
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

	// Update category if provided
	if data.CategoryID != nil && *data.CategoryID > 0 {
		if err := repo.storage.UpsertProductCategory(ctx, productID, *data.CategoryID); err != nil {
			return err
		}
	}

	return nil
}
