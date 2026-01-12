package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

// ListAllCategories returns all categories for tree building
func (s *sqlStore) ListAllCategories(
	ctx context.Context,
) ([]productmodel.CategoryListItem, error) {
	var result []productmodel.CategoryListItem
	if err := s.db.Table(productmodel.Category{}.TableName()).
		Select("id, name, NULLIF(parent_id, 0) as parent_id").
		Where("status = 1").
		Order("parent_id, id").
		Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
