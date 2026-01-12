package productstorage

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

// GetVariantWithAttributes returns a variant with full attribute info
func (s *sqlStore) GetVariantWithAttributes(
	ctx context.Context,
	variantID int,
	productID int,
) ([]productmodel.VariantAttrFullRow, error) {
	var rows []productmodel.VariantAttrFullRow
	if err := s.db.Table(productmodel.Variant{}.TableName()+" v").
		Select(`
			v.id AS variant_id,
			v.sku,
			v.price,
			v.stock_quantity,
			v.status,
			v.created_at,
			v.updated_at,
			a.id AS attribute_id,
			a.name AS attribute_name,
			av.id AS attribute_value_id,
			av.value AS attribute_value
		`).
		Joins("LEFT JOIN variant_attribute_values vav ON v.id = vav.variant_id").
		Joins("LEFT JOIN attribute_values av ON av.id = vav.attribute_value_id").
		Joins("LEFT JOIN attributes a ON a.id = av.attribute_id").
		Where("v.id = ? AND v.product_id = ? AND v.status = ?", variantID, productID, common.SystemStatusActive).
		Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
