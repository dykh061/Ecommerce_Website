package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

// TxStore là interface transaction-level
// Storage tự định nghĩa, không phụ thuộc repository
type TxStore interface {
	CreateVariant(ctx context.Context, data *productmodel.VariantCreate) error
	CreateVariantAttributeValues(
		ctx context.Context,
		data []productmodel.VariantAttributeValue,
	) error

	AdjustVariantStock(
		ctx context.Context,
		varianId int,
		by int,
	) error
	FindVariantWithAtributesValue(
		cxt context.Context,
		productId int,
		attributeValueIds []int,
	) (*productmodel.Variant, error)
}
