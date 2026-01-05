package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) UpdateVariant(
	ctx context.Context,
	condition map[string]interface{},
	data *productmodel.VariantUpdate,
) error {
	db := s.db.WithContext(ctx).
		Model(&productmodel.Variant{}).
		Where(condition).
		Updates(data)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
