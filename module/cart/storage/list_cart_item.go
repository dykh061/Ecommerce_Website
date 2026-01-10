package cartstorage

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

func (s *sqlStore) ListCartItems(
	ctx context.Context,
	cartId int,
) ([]cartmodel.CartItem, error) {
	var items []cartmodel.CartItem
	if err := s.db.
		Where("cart_id = ?", cartId).
		Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
