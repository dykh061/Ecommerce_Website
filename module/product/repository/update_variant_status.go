package productrepository

import "context"

type UpdateVariantStatusStorage interface {
	UpdateVariantStatus(ctx context.Context, variantID int, productID int, status int) error
}

type updateVariantStatusRepo struct {
	storage UpdateVariantStatusStorage
}

func NewUpdateVariantStatusRepo(storage UpdateVariantStatusStorage) *updateVariantStatusRepo {
	return &updateVariantStatusRepo{storage: storage}
}

func (repo *updateVariantStatusRepo) UpdateStatus(
	ctx context.Context,
	variantID int,
	productID int,
	status int,
) error {
	return repo.storage.UpdateVariantStatus(ctx, variantID, productID, status)
}
