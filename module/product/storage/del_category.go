package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) DeleteCategory(
	ctx context.Context,
	id, parentId int,
) error {
	if err := s.db.Table(productmodel.Category{}.TableName()).
		Where("id = ? AND parent_id = ? ", id, parentId).
		Delete(nil).
		Error; err != nil {
		return err
	}
	return nil
}
