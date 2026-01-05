package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) CreateProductGallery(
	ctx context.Context,
	data *productmodel.GalleryCreate,
) error {
	return s.db.Create(data).Error
}
