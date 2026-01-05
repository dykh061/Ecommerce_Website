package sellerbusiness

import (
	"OpenMarket/common"
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type ListSellerRepo interface {
	GetListSeller(
		ctx context.Context,
		filter *sellermodel.SellerFilter,
		paging *common.Paging,
	) ([]sellermodel.Seller, error)
}

type listSellerBusiness struct {
	repo ListSellerRepo
}

func NewListSellerBusiness(repo ListSellerRepo) *listSellerBusiness {
	return &listSellerBusiness{repo: repo}
}

func (biz *listSellerBusiness) ListSellers(
	ctx context.Context,
	filter *sellermodel.SellerFilter,
	paging *common.Paging,
	morekeys ...string,
) ([]sellermodel.Seller, error) {
	result, err := biz.repo.GetListSeller(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrorDB(err)
	}
	if result == nil {
		return []sellermodel.Seller{}, nil
	}
	return result, nil
}
