package cartstorage

import (
	"OpenMarket/common"
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

func (s *sqlStore) ListCartItemsDetail(
	ctx context.Context,
	cartId int,
) ([]cartmodel.CartItemDetailRow, error) {
	var rows []cartmodel.CartItemDetailRow
	if err := s.db.WithContext(ctx).
		Table(cartmodel.CartItem{}.TableName() + " ci").
		Select(`
			ci.variant_id AS variant_id,
			ci.quantity AS quantity,
			v.price AS price,
			v.stock_quantity AS stock_quantity,
			p.id AS product_id,
			p.name AS product_name,
			g.image_url AS image_url,
			a.name AS attribute_name,
			av.value AS attribute_value
		`).
		Joins("JOIN variants v ON v.id = ci.variant_id").
		Joins("JOIN products p ON p.id = v.product_id").
		Joins("LEFT JOIN galleries g ON g.product_id = p.id AND g.is_main = 1").
		Joins("LEFT JOIN variant_attribute_values vav ON vav.variant_id = v.id").
		Joins("LEFT JOIN attribute_values av ON av.id = vav.attribute_value_id").
		Joins("LEFT JOIN attributes a ON a.id = av.attribute_id").
		Where("ci.cart_id = ?", cartId).
		Where("v.status = ?", common.SystemStatusActive).
		Where("p.status = ?", common.SystemStatusActive).
		Order("ci.id asc").
		Order("a.id asc").
		Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
