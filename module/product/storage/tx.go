package productstorage

import (
	productmodel "OpenMarket/module/product/model"
	"context"

	"github.com/shopspring/decimal"
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
	UpdateVariant(
		ctx context.Context,
		condition map[string]interface{},
		upprice *decimal.Decimal,
	) error
}
