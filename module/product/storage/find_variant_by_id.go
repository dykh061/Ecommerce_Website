package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) FindVariantByID(
	ctx context.Context,
	id int,
) (*productmodel.Variant, error) {
	var data productmodel.Variant
	if err := s.db.Table(productmodel.Variant{}.TableName()).
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil

}
