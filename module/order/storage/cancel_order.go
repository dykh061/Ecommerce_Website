package orderstorage

import (
	"OpenMarket/common"
	ordermodel "OpenMarket/module/order/model"
	"context"
	"errors"
)

func (s *sqlStore) CancelOrder(
	ctx context.Context,
	id int,
	reason string,
) error {
	tx := s.db.Table(ordermodel.Order{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":        ordermodel.OrderCancelled,
			"cancel_reason": reason,
		})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return common.ErrEntityNotFound("Order", errors.New("record not found"))
	}

	return nil
}
