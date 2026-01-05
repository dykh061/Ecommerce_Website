package productstorage

import "context"

func (s *sqlStore) WithTransaction(
	ctx context.Context,
	fn func(tx TxStore) error,
) error {
	tx := s.db.WithContext(ctx).Begin() // mở transaction
	if tx.Error != nil {
		return tx.Error
	}

	txStore := &sqlStore{db: tx} // tạo một sqlStore mới với transaction

	// gọi hàm truyền vào với txStore
	if err := fn(txStore); err != nil {
		tx.Rollback()
		return err
	}

	// commit transaction nếu không có lỗi
	return tx.Commit().Error
}
