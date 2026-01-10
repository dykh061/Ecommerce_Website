package cartstorage

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

func (s *sqlStore) FindCart(
	ctx context.Context,
	userId int,
) (*cartmodel.Cart, error) {
	var data cartmodel.Cart
	if err := s.db.Table(cartmodel.Cart{}.TableName()).
		Where("user_id = ?", userId).
		First(&data).
		Error; err != nil {
		return nil, err
	}
	return &data, nil
}
