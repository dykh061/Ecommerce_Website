package productbusiness

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type listVariantBusiness struct {
	repo ListVariantRepo
}

func NewListVariantBusiness(repo ListVariantRepo) *listVariantBusiness {
	return &listVariantBusiness{repo: repo}
}

func (business *listVariantBusiness) ListVariant(
	ctx context.Context,
	productID int,
) ([]productmodel.VariantDetail, error) {
	return business.repo.ListVariant(ctx, productID)
}
