package cartstorage

import (
	"OpenMarket/common"
	cartmodel "OpenMarket/module/cart/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlStore) UpdateCartItem(
	ctx context.Context,
	id, by int,
) error {
	db := s.db.
		Table(cartmodel.CartItem{}.TableName()).
		Where("id = ? AND quantity + ? >= 0 ", id, by).
		UpdateColumn("quantity", gorm.Expr("quantity + ?", by))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return common.InvalidRequestError(errors.New("not enough quantity or item not found"))
	}
	return nil
}
