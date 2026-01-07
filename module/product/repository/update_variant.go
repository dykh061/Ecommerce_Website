package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	productstorage "OpenMarket/module/product/storage"
	"context"
)

type TransactionProvider interface {
	WithTransaction(
		ctx context.Context,
		fn func(tx productstorage.TxStore) error,
	) error
}

type updateVariantRepo struct {
	txProvider TransactionProvider
}

func NewUpdateVariantRepo(tp TransactionProvider) *updateVariantRepo {
	return &updateVariantRepo{txProvider: tp}
}

func (repo *updateVariantRepo) UpdateVariant(
	ctx context.Context,
	variantId int,
	condition map[string]interface{},
	data *productmodel.VariantUpdate,
) error {
	return repo.txProvider.WithTransaction(ctx, func(tx productstorage.TxStore) error {
		if data.Price != nil {
			if err := tx.UpdateVariant(ctx, condition, data.Price); err != nil {
				return err
			}
		}
		if data.StockQuantity != nil {
			if err := tx.AdjustVariantStock(ctx, variantId, *data.StockQuantity); err != nil {
				return err
			}
		}
		return nil
	})
}
