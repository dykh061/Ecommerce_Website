package cartstorage

import (
	cartrepository "OpenMarket/module/cart/repository"
	"context"
)

func (s *sqlStore) WithTransaction(
	ctx context.Context,
	fn func(tx cartrepository.TxStore) error,
) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	txStore := &sqlStore{db: tx}

	if err := fn(txStore); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
