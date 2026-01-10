package productstorage

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) DeleteProduct(
	ctx context.Context,
	productId int,
) error {
	if err := s.db.Table(productmodel.Product{}.TableName()).
		Where("id = ?", productId).
		Updates(map[string]interface{}{
			"status": common.SystemStatusDeleted,
		}).Error; err != nil {
		return err
	}
	return nil
}
