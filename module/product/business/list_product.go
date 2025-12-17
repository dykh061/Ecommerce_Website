package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

type ListProductStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *productmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]productmodel.Product, error)
}

type listProductBusiness struct {
	store ListProductStore
}

func NewListProductBusiness(store ListProductStore) *listProductBusiness {
	return &listProductBusiness{store: store}
}

func (biz *listProductBusiness) ListProduct(
	ctx context.Context,
	filter *productmodel.Filter,
	paging *common.Paging,
) ([]productmodel.Product, error) {

	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
