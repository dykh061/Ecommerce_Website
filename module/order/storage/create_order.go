package orderstorage

import (
	ordermodel "OpenMarket/module/order/model"
	"context"
)

func (s *sqlStore) CreateOrder(
	ctx context.Context,
	data *ordermodel.Order,
) error {
	return s.db.Table(ordermodel.Order{}.TableName()).Create(data).Error
}
