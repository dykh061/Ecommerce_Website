package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) CreateVariant(ctx context.Context, data *productmodel.VariantCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
