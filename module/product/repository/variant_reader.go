package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type VariantReader interface {
	FindVariantByID(
		ctx context.Context,
		id int,
	) (*productmodel.Variant, error)
}
