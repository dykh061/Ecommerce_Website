package orderstorage

import (
	ordermodel "OpenMarket/module/order/model"
	"context"
)

func (s *sqlStore) FindOrderByID(
	ctx context.Context,
	id int,
) (*ordermodel.Order, error) {
	var result ordermodel.Order
	if err := s.db.Table(ordermodel.Order{}.TableName()).
		Where("id = ?", id).
		First(&result).
		Error; err != nil {
		return nil, err
	}
	return &result, nil
}
