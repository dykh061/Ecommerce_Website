package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

// GetCategoryAttributes returns attributes and their values for a category
func (s *sqlStore) GetCategoryAttributes(
	ctx context.Context,
	categoryID int,
) ([]productmodel.CategoryAttributeRow, error) {
	var rows []productmodel.CategoryAttributeRow
	if err := s.db.Table("category_attributes ca").
		Select(`
			a.id AS attribute_id,
			a.name AS attribute_name,
			av.id AS value_id,
			av.value AS attribute_value
		`).
		Joins("JOIN attributes a ON a.id = ca.attribute_id").
		Joins("LEFT JOIN attribute_values av ON av.attribute_id = a.id").
		Where("ca.category_id = ?", categoryID).
		Order("a.id, av.id").
		Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

// CategoryExists checks if a category exists
func (s *sqlStore) CategoryExists(
	ctx context.Context,
	categoryID int,
) (bool, error) {
	var count int64
	if err := s.db.Table(productmodel.Category{}.TableName()).
		Where("id = ? AND status = 1", categoryID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
