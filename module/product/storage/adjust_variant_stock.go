package productstorage

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlStore) AdjustVariantStock(
	ctx context.Context,
	varianId int,
	by int,
) error {
	if by == 0 {
		return nil
	}
	db := s.db.Table(productmodel.Variant{}.TableName()).
		Where("id = ? AND stock_quantity + ? >= 0", varianId, by).
		UpdateColumn("stock_quantity", gorm.Expr("stock_quantity + ?", by))

	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		// không có rows bị ảnh hưởng => có thể vì id không tồn tại hoặc không đủ stock
		return common.InvalidRequestError(errors.New("not enough stock or variant not found"))
	}

	return nil
}
