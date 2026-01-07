package productstorage

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) ListVariant(
	ctx context.Context,
	productID int,
) ([]productmodel.VariantAttrRow, error) {
	var rows []productmodel.VariantAttrRow
	if err := s.db.Table(productmodel.Variant{}.TableName()+" v").
		Select(`
			v.id AS variant_id,
			v.sku,
			v.price,
			v.stock_quantity,
			a.id AS attribute_id,
			a.name AS attribute_name,
			av.value AS attribute_value
		`).
		Joins(`join variant_attribute_values vav on v.id = vav.variant_id`).
		Joins("JOIN attribute_values av ON av.id = vav.attribute_value_id").
		Joins("JOIN attributes a ON a.id = av.attribute_id").
		Where("v.product_id = ? and v.status = ?", productID, common.SystemStatusActive).
		Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
