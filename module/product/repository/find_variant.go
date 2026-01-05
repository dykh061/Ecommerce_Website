package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type FindVariantStorage interface {
	FindVariantWithAtributesValue(
		cxt context.Context,
		productId int,
		attributeValueIds []int,
	) (*productmodel.Variant, error)
}

type findVariantRepo struct {
	storage FindVariantStorage
}

func NewFindVariantRepo(storage FindVariantStorage) *findVariantRepo {
	return &findVariantRepo{storage: storage}
}

func (repo *findVariantRepo) FindVariantByAttributesvalue(
	ctx context.Context,
	productId int,
	attributeValueIds []int,
) (*productmodel.Variant, error) {
	return repo.storage.FindVariantWithAtributesValue(
		ctx,
		productId,
		attributeValueIds,
	)
}
