package cartstorage

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

func (s *sqlStore) DeleteCartItem(
	ctx context.Context,
	id int,
) error {
	return s.db.
		Table(cartmodel.CartItem{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error
}
