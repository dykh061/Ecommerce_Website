package cartstorage

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

func (s *sqlStore) FindCartItem(
	ctx context.Context,
	cartId, variantId int,
) (*cartmodel.CartItem, error) {
	var item cartmodel.CartItem
	err := s.db.
		Where("cart_id = ? AND variant_id = ?", cartId, variantId).
		First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}
