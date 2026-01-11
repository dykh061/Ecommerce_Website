package orderbusiness

import (
	"OpenMarket/common"
	ordermodel "OpenMarket/module/order/model"
	"context"
)

type GetListOrderStore interface {
	ListOrders(
		ctx context.Context,
		filter *ordermodel.FilterOrder,
		paging *common.Paging,
	) ([]ordermodel.Order, error)
}

type getListOrderBiz struct {
	store GetListOrderStore
}

func NewGetListOrderBiz(store GetListOrderStore) *getListOrderBiz {
	return &getListOrderBiz{store: store}
}

func (biz *getListOrderBiz) GetListOrder(
	ctx context.Context,
	filter *ordermodel.FilterOrder,
	paging *common.Paging,
) ([]ordermodel.Order, error) {
	return biz.store.ListOrders(ctx, filter, paging)
}
