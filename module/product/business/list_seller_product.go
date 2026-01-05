package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type ListSellerProductRepo interface {
	ListPublicProducts(
		ctx context.Context,
		filter *productmodel.Filter,
		paging *common.Paging,
	) ([]productmodel.ProductListItem, error)
}

type listSellerProductBusiness struct {
	repo   ListSellerProductRepo
	finder FindSellerByID
}

func NewListSellerProductBusiness(
	repo ListSellerProductRepo,
	finder FindSellerByID,
) *listSellerProductBusiness {
	return &listSellerProductBusiness{repo: repo, finder: finder}
}

func (biz *listSellerProductBusiness) ListSellerProducts(
	ctx context.Context,
	userId int,
	filter *productmodel.Filter,
	paging *common.Paging,
) ([]productmodel.ProductListItem, error) {
	seller, err := biz.finder.FindActiveSellerWithUserID(ctx, userId)
	if err != nil {
		return nil, common.ErrEntityNotFound(sellermodel.EntityName, err)
	}
	filter.SellerID = &seller.Id

	result, err := biz.repo.ListPublicProducts(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
