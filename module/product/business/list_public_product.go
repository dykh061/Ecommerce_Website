package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

type ListPublicProductRepo interface {
	ListPublicProducts(
		ctx context.Context,
		filter *productmodel.Filter,
		paging *common.Paging,
	) ([]productmodel.ProductListItem, error)
}

type listPublicProductBusiness struct {
	repo ListPublicProductRepo
}

func NewListPublicProductBusiness(repo ListPublicProductRepo) *listPublicProductBusiness {
	return &listPublicProductBusiness{repo: repo}
}

func (business *listPublicProductBusiness) ListPublicProducts(
	ctx context.Context,
	filter *productmodel.Filter,
	paging *common.Paging,
) ([]productmodel.ProductListItem, error) {
	return business.repo.ListPublicProducts(ctx, filter, paging)
}
