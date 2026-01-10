package productstorage

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) DeleteVariant(
	ctx context.Context,
	variantId int,
) error {
	if err := s.db.Table(productmodel.Variant{}.TableName()).
		Where("id = ?", variantId).
		Updates(map[string]interface{}{
			"status": common.SystemStatusDeleted,
		}).Error; err != nil {
		return err
	}
	return nil
}
