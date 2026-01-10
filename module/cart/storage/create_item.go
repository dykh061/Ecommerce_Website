package cartstorage

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

func (s *sqlStore) CreateItem(
	ctx context.Context,
	item *cartmodel.CartItemCreate,
) error {
	return s.db.Create(item).Error
}
