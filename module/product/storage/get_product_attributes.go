package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

// GetProductAttributes returns all attributes and their values for variants of a product
func (s *sqlStore) GetProductAttributes(
	ctx context.Context,
	productID int,
) ([]productmodel.AttributeValueRow, error) {
	var rows []productmodel.AttributeValueRow
	if err := s.db.Table("attributes a").
		Select(`
			a.id AS attribute_id,
			a.name AS attribute_name,
			av.id AS value_id,
			av.value AS attribute_value
		`).
		Joins("JOIN attribute_values av ON av.attribute_id = a.id").
		Joins("JOIN variant_attribute_values vav ON vav.attribute_value_id = av.id").
		Joins("JOIN variants v ON v.id = vav.variant_id").
		Where("v.product_id = ?", productID).
		Group("a.id, a.name, av.id, av.value").
		Order("a.id, av.id").
		Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

// GetProductCategoryID returns the category_id for a product
func (s *sqlStore) GetProductCategoryID(
	ctx context.Context,
	productID int,
) (*int, error) {
	var categoryID *int
	if err := s.db.Table(productmodel.ProductCategory{}.TableName()).
		Select("category_id").
		Where("product_id = ?", productID).
		Limit(1).
		Scan(&categoryID).Error; err != nil {
		return nil, err
	}
	return categoryID, nil
}
