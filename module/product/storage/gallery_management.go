package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

// SetMainGallery sets a gallery as main and unsets others
func (s *sqlStore) SetMainGallery(
	ctx context.Context,
	productID int,
	galleryID int,
) error {
	// Start transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Reset all galleries for this product
	if err := tx.Table(productmodel.Gallery{}.TableName()).
		Where("product_id = ?", productID).
		Update("is_main", false).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Set the specified gallery as main
	if err := tx.Table(productmodel.Gallery{}.TableName()).
		Where("id = ? AND product_id = ?", galleryID, productID).
		Update("is_main", true).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DeleteGallery deletes a gallery by ID
func (s *sqlStore) DeleteGallery(
	ctx context.Context,
	productID int,
	galleryID int,
) error {
	return s.db.Table(productmodel.Gallery{}.TableName()).
		Where("id = ? AND product_id = ?", galleryID, productID).
		Delete(&productmodel.Gallery{}).Error
}

// GalleryExists checks if a gallery exists and belongs to a product
func (s *sqlStore) GalleryExists(
	ctx context.Context,
	productID int,
	galleryID int,
) (bool, error) {
	var count int64
	if err := s.db.Table(productmodel.Gallery{}.TableName()).
		Where("id = ? AND product_id = ?", galleryID, productID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
