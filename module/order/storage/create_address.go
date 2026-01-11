package orderstorage

import (
	ordermodel "OpenMarket/module/order/model"
	"context"
)

func (s *sqlStore) CreateAddress(
	ctx context.Context,
	data *ordermodel.OrderAddressCreate,
) error {
	return s.db.Table(ordermodel.OrderAddressCreate{}.TableName()).Create(data).Error
}
