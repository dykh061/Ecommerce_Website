package orderstorage

import (
	ordermodel "OpenMarket/module/order/model"
	"context"
)

func (s *sqlStore) ListOrderItemsByOrderID(
	ctx context.Context,
	orderID int,
) ([]ordermodel.OrderItem, error) {
	var items []ordermodel.OrderItem
	if err := s.db.Table(ordermodel.OrderItem{}.TableName()).
		Where("order_id = ?", orderID).
		Find(&items).
		Error; err != nil {
		return nil, err
	}
	return items, nil
}
