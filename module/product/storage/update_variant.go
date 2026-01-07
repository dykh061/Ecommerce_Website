package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"

	"github.com/shopspring/decimal"
)

func (s *sqlStore) UpdateVariant(
	ctx context.Context,
	condition map[string]interface{},
	upprice *decimal.Decimal,
) error {

	db := s.db.WithContext(ctx).
		Model(&productmodel.Variant{}).
		Where(condition).
		Updates(map[string]interface{}{
			"price": upprice,
		})
	if db.Error != nil {
		return db.Error
	}
	return nil
}
