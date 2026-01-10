package cartstorage

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

func (s *sqlStore) CreateCart(
	ctx context.Context,
	data *cartmodel.CartCreate,
) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
