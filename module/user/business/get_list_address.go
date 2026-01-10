package userbusiness

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

type GetListAddressStore interface {
	GetListAddress(
		ctx context.Context,
		userId int,
	) ([]usermodel.UserAddress, error)
}
type getListAddressBusiness struct {
	store GetListAddressStore
}

func NewGetListAddressBusiness(store GetListAddressStore) *getListAddressBusiness {
	return &getListAddressBusiness{store: store}
}

func (biz *getListAddressBusiness) GetListAddress(
	ctx context.Context,
	userId int,
) ([]usermodel.UserAddress, error) {
	return biz.store.GetListAddress(ctx, userId)
}
