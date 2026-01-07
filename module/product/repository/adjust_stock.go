package productrepository

import "context"

type AdjustStockStorage interface {
	AdjustVariantStock(
		ctx context.Context,
		varianId int,
		by int,
	) error
}

type adjustStockStorage struct {
	storage AdjustStockStorage
}

func NewAdjustStockRepo(
	storage AdjustStockStorage,
) *adjustStockStorage {
	return &adjustStockStorage{
		storage: storage,
	}
}

func (repo *adjustStockStorage) AdjustStock(
	ctx context.Context,
	variantId int,
	by int,
) error {
	return repo.storage.AdjustVariantStock(ctx, variantId, by)
}
