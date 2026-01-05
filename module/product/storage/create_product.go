package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *productmodel.ProductCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
