package sellerrepository

import (
	"OpenMarket/common"
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type ListSellerStorage interface {
	ListSellers(
		ctx context.Context,
		filter *sellermodel.SellerFilter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]sellermodel.Seller, error)
}

type listSellerRepo struct {
	storage ListSellerStorage
}

func NewListSellerRepo(storage ListSellerStorage) *listSellerRepo {
	return &listSellerRepo{storage: storage}
}

func (repo *listSellerRepo) GetListSeller(
	ctx context.Context,
	filter *sellermodel.SellerFilter,
	paging *common.Paging,
) ([]sellermodel.Seller, error) {
	return repo.storage.ListSellers(ctx, filter, paging)
}
