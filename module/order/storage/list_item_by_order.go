package orderstorage

import (
	ordermodel "OpenMarket/module/order/model"
	"context"
)

func (s *sqlStore) ListOrderItemsByOrderID(
	ctx context.Context,
	orderID int,
) ([]ordermodel.OrderItemDetailRow, error) {
	var rows []ordermodel.OrderItemDetailRow
	if err := s.db.WithContext(ctx).
		Table(ordermodel.OrderItem{}.TableName()+" oi").
		Select(`
			oi.id AS order_item_id,
			oi.variant_id AS variant_id,
			oi.quantity AS quantity,
			oi.price AS price,
			p.name AS product_name,
			g.image_url AS image_url,
			a.name AS attribute_name,
			av.value AS attribute_value
		`).
		Joins("LEFT JOIN variants v ON v.id = oi.variant_id").
		Joins("LEFT JOIN products p ON p.id = v.product_id").
		Joins("LEFT JOIN galleries g ON g.product_id = p.id AND g.is_main = 1").
		Joins("LEFT JOIN variant_attribute_values vav ON vav.variant_id = v.id").
		Joins("LEFT JOIN attribute_values av ON av.id = vav.attribute_value_id").
		Joins("LEFT JOIN attributes a ON a.id = av.attribute_id").
		Where("oi.order_id = ?", orderID).
		Order("oi.id asc").
		Order("a.id asc").
		Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
