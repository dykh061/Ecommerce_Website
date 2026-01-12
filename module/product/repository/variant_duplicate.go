package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type VariantDuplicateStorage interface {
	FindActiveVariantWithAttributes(
		ctx context.Context,
		productID int,
		attributeValueIDs []int,
		excludeVariantID *int,
	) (*productmodel.Variant, error)
}

type variantDuplicateRepo struct {
	storage VariantDuplicateStorage
}

func NewVariantDuplicateRepo(storage VariantDuplicateStorage) *variantDuplicateRepo {
	return &variantDuplicateRepo{storage: storage}
}

func (repo *variantDuplicateRepo) CheckDuplicate(
	ctx context.Context,
	productID int,
	attributeValueIDs []int,
	excludeVariantID *int,
) (bool, error) {
	variant, err := repo.storage.FindActiveVariantWithAttributes(ctx, productID, attributeValueIDs, excludeVariantID)
	if err != nil {
		// If record not found, it's not a duplicate
		return false, nil
	}
	return variant != nil && variant.Id > 0, nil
}
