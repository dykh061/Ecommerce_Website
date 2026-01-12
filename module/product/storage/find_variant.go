package productstorage

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

func (s *sqlStore) FindVariantWithAtributesValue(
	cxt context.Context,
	productId int,
	attributeValueIds []int,
) (*productmodel.Variant, error) {
	var variant productmodel.Variant
	tableName := productmodel.Variant{}.TableName()
	err := s.db.
		Table(tableName).
		Joins("join variant_attribute_values vav on vav.variant_id = "+tableName+".id").
		Where(tableName+".product_id = ?", productId).
		Where(tableName+".status = ?", common.SystemStatusActive).
		Where("vav.attribute_value_id IN ?", attributeValueIds).
		Group(tableName+".id").
		Having("COUNT(DISTINCT vav.attribute_value_id) = ?", len(attributeValueIds)).
		First(&variant).Error

	if err != nil {
		return nil, err
	}
	return &variant, nil
}
