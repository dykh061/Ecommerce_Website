package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"

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
		Where("id = ?", varianId).
		UpdateColumn("stock_quantity", gorm.Expr("stock_quantity + ?", by))
	return db.Error
}
