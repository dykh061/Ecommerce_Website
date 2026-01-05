package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) CreateVariantAttributeValues(
	ctx context.Context,
	data []productmodel.VariantAttributeValue,
) error {
	if len(data) == 0 {
		return nil
	}
	return s.db.Create(&data).Error
}
