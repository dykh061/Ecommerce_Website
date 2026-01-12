package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

// UpsertProductCategory creates or updates the product-category relationship
func (s *sqlStore) UpsertProductCategory(
	ctx context.Context,
	productID int,
	categoryID int,
) error {
	// First, delete existing relationship
	if err := s.db.WithContext(ctx).
		Table(productmodel.ProductCategory{}.TableName()).
		Where("product_id = ?", productID).
		Delete(&productmodel.ProductCategory{}).Error; err != nil {
		return err
	}

	// Then insert new relationship
	pc := productmodel.ProductCategory{
		ProductId:  productID,
		CategoryId: categoryID,
	}
	return s.db.WithContext(ctx).Create(&pc).Error
}

// DeleteProductCategory removes the product-category relationship
func (s *sqlStore) DeleteProductCategory(
	ctx context.Context,
	productID int,
) error {
	return s.db.WithContext(ctx).
		Table(productmodel.ProductCategory{}.TableName()).
		Where("product_id = ?", productID).
		Delete(&productmodel.ProductCategory{}).Error
}
