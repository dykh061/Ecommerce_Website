package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) FindProductCategory(
	ctx context.Context,
	id, parentId int,
) error {
	var result productmodel.Category
	if err := s.db.Table(productmodel.Category{}.TableName()).
		Where("id = ? AND parent_id = ?", id, parentId).
		Find(&result).
		Error; err != nil {
		return err
	}
	return nil
}
