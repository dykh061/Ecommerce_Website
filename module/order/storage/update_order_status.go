package orderstorage

import (
	"OpenMarket/common"
	ordermodel "OpenMarket/module/order/model"
	"context"
	"errors"
)

func (s *sqlStore) UpdateOrderStatus(
	ctx context.Context,
	id int,
	status string,
) error {
	tx := s.db.Table(ordermodel.OrderStatusUpdate{}.TableName()).
		Where("id = ?", id).
		Update("status", status)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return common.ErrEntityNotFound("Order", errors.New("Record not found"))
	}
	return nil
}
