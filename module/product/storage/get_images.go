package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) GetImages(
	context context.Context,
	productID int,
) ([]string, error) {
	var urls []string
	if err := s.db.Table(productmodel.Gallery{}.TableName()).
		Where("product_id = ?", productID).
		Order("is_main desc, id asc").
		Pluck("image_url", &urls).
		Error; err != nil {
		return nil, err
	}
	return urls, nil
}
