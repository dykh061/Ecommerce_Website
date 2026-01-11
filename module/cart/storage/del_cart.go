package cartstorage

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

func (s *sqlStore) DeleteCart(
	ctx context.Context,
	userId int,
) error {
	return s.db.Table(cartmodel.Cart{}.TableName()).Where("user_id = ?", userId).Delete(nil).Error
}
