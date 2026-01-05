package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type UpdateVariantStorage interface {
	UpdateVariant(
		ctx context.Context,
		condition map[string]interface{},
		data *productmodel.VariantUpdate,
	) error
}

type updateVariantRepo struct {
	storage UpdateVariantStorage
}

func NewUpdateVariantRepo(storage UpdateVariantStorage) *updateVariantRepo {
	return &updateVariantRepo{storage: storage}
}

func (repo *updateVariantRepo) UpdateVariant(
	ctx context.Context,
	condition map[string]interface{},
	data *productmodel.VariantUpdate,
) error {
	return repo.storage.UpdateVariant(ctx, condition, data)
}
