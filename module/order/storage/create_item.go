package orderstorage

import (
	ordermodel "OpenMarket/module/order/model"
	"context"
)

func (s *sqlStore) CreateOrderItem(
	ctx context.Context,
	item *ordermodel.OrderItemCreate,
) error {
	return s.db.Table(ordermodel.OrderItem{}.TableName()).Create(item).Error
}
