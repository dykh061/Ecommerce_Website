package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) UpdateProduct(
	ctx context.Context,
	condition map[string]interface{},
	data *productmodel.ProductUpdate,
) error {
	if err := s.db.WithContext(ctx).
		Table(productmodel.ProductUpdate{}.TableName()).
		Where(condition).
		Updates(data).
		Error; err != nil {
		return err
	}
	return nil
}
