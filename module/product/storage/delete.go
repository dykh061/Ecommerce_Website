package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(productmodel.ProductUpdate{}.TableName()).
		Where("id=?", id).
		Updates(map[string]interface{}{"status": "deleted"}).
		Error; err != nil {
		return err
	}
	return nil
}
