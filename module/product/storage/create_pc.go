package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) AddProductCategory(
	ctx context.Context,
	pc *productmodel.ProductCategory,
) error {
	if err := s.db.Create(&pc).Error; err != nil {
		return err
	}
	return nil
}
