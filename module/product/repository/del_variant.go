package productrepository

import "context"

type DeleteVariantStorage interface {
	DeleteVariant(
		ctx context.Context,
		variantId int,
	) error
}

type deleteVariantRepo struct {
	storage DeleteVariantStorage
}

func NewDeleteVariantRepo(
	storage DeleteVariantStorage,
) *deleteVariantRepo {
	return &deleteVariantRepo{storage: storage}
}

func (repo *deleteVariantRepo) DeleteVariant(
	ctx context.Context,
	variantId int,
) error {
	return repo.storage.DeleteVariant(ctx, variantId)
}
