package productbusiness

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type ListVariantRepo interface {
	ListVariant(
		ctx context.Context,
		productID int,
	) ([]productmodel.VariantDetail, error)
}
