package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

// GetGalleries returns gallery items for a product
func (s *sqlStore) GetGalleries(
	ctx context.Context,
	productID int,
) ([]productmodel.GalleryItem, error) {
	var galleries []productmodel.GalleryItem
	if err := s.db.Table(productmodel.Gallery{}.TableName()).
		Select("id, image_url as url, is_main").
		Where("product_id = ?", productID).
		Order("is_main desc, id asc").
		Scan(&galleries).
		Error; err != nil {
		return nil, err
	}
	return galleries, nil
}
