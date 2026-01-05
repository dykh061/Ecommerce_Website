package productrepository

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

type ListPublicProductStorage interface {
	ListProduct(
		ctx context.Context,
		filter *productmodel.Filter,
		paging *common.Paging,
	) ([]productmodel.ProductListItem, error)
}
type listPublicProductRepo struct {
	storage ListPublicProductStorage
}

func NewListPublicProductRepo(storage ListPublicProductStorage) *listPublicProductRepo {
	return &listPublicProductRepo{storage: storage}
}

func (repo *listPublicProductRepo) ListPublicProducts(
	ctx context.Context,
	filter *productmodel.Filter,
	paging *common.Paging,
) ([]productmodel.ProductListItem, error) {
	return repo.storage.ListProduct(ctx, filter, paging)
}
