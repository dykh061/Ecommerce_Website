package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) DeleteProductCategoryByProductID(
	ctx context.Context,
	productID int,
) error {
	if err := s.db.Table(productmodel.ProductCategory{}.TableName()).
		Where("product_id = ?", productID).
		Delete(nil).
		Error; err != nil {
		return err
	}
	return nil
}
