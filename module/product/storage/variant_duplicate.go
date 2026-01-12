package productstorage

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

// FindActiveVariantWithAttributes finds an active variant with exact attribute_value_ids
// Returns nil if not found
func (s *sqlStore) FindActiveVariantWithAttributes(
	ctx context.Context,
	productID int,
	attributeValueIDs []int,
	excludeVariantID *int,
) (*productmodel.Variant, error) {
	if len(attributeValueIDs) == 0 {
		return nil, nil
	}

	tableName := productmodel.Variant{}.TableName()
	query := s.db.
		Table(tableName).
		Joins("JOIN variant_attribute_values vav ON vav.variant_id = "+tableName+".id").
		Where(tableName+".product_id = ?", productID).
		Where(tableName+".status = ?", common.SystemStatusActive).
		Where("vav.attribute_value_id IN ?", attributeValueIDs).
		Group(tableName + ".id").
		Having("COUNT(DISTINCT vav.attribute_value_id) = ?", len(attributeValueIDs))

	if excludeVariantID != nil && *excludeVariantID > 0 {
		query = query.Where(tableName+".id != ?", *excludeVariantID)
	}

	var variant productmodel.Variant
	if err := query.First(&variant).Error; err != nil {
		return nil, err
	}
	return &variant, nil
}

// UpdateVariantStatus updates the status of a variant
func (s *sqlStore) UpdateVariantStatus(
	ctx context.Context,
	variantID int,
	productID int,
	status int,
) error {
	return s.db.Table(productmodel.Variant{}.TableName()).
		Where("id = ? AND product_id = ?", variantID, productID).
		Update("status", status).Error
}
