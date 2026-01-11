package orderstorage

import (
	ordermodel "OpenMarket/module/order/model"
	"context"

	"github.com/shopspring/decimal"
)

func (s *sqlStore) UpTotalAmount(
	ctx context.Context,
	totalAmount decimal.Decimal,
	id int,
) error {
	return s.db.Table(ordermodel.Order{}.TableName()).Where("id = ?", id).Update("total_amount", totalAmount).Error
}
