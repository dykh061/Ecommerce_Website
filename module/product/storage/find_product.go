package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) FindProductWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*productmodel.Product, error) {
	var data productmodel.Product
	if err := s.db.Where(condition).
		First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
