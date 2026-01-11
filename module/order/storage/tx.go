package orderstorage

import (
	cartstorage "OpenMarket/module/cart/storage"
	orderrepository "OpenMarket/module/order/repository"
	productstorage "OpenMarket/module/product/storage"
	"context"
)

func (s *sqlStore) WithTransaction(
	ctx context.Context,
	fn func(tx orderrepository.TxStore) error,
) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	orderStore := &sqlStore{db: tx}
	productStore := productstorage.NewSQLStore(tx)
	cartStore := cartstorage.NewSQLStore(tx)

	txStore := &txStore{
		db:            tx,
		orderStore:    orderStore,
		stockStore:    productStore, // implements AdjustStockStorage
		variantReader: productStore, // implements VariantReader
		cartStore:     cartStore,
	}

	if err := fn(txStore); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
