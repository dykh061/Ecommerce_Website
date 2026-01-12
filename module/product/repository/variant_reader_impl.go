package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type VariantReaderStorage interface {
	FindVariantByID(ctx context.Context, id int) (*productmodel.Variant, error)
}

type variantReaderRepo struct {
	storage VariantReaderStorage
}

func NewVariantReaderRepo(storage VariantReaderStorage) *variantReaderRepo {
	return &variantReaderRepo{storage: storage}
}

func (repo *variantReaderRepo) FindVariantByID(
	ctx context.Context,
	id int,
) (*productmodel.Variant, error) {
	return repo.storage.FindVariantByID(ctx, id)
}
