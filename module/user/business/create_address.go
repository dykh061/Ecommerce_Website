package userbusiness

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
)

type CreateAddressStore interface {
	CreateAddress(
		ctx context.Context,
		create *usermodel.UserAddressCreate,
	) error
}

type createAddressBusiness struct {
	store CreateAddressStore
}

func NewCreateAddressBusiness(store CreateAddressStore) *createAddressBusiness {
	return &createAddressBusiness{store: store}
}

func (biz *createAddressBusiness) CreateAddress(
	ctx context.Context,
	userId int,
	create *usermodel.UserAddressCreate,
) error {
	create.UserId = userId
	if err := biz.store.CreateAddress(ctx, create); err != nil {
		return common.ErrCannotCreateEntity("User Address", err)
	}
	return nil
}
