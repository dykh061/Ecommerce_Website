package userbusiness

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
)

type UpdateAddressStore interface {
	UpdateAddress(
		ctx context.Context,
		userId, id int,
		data *usermodel.UserAddressUpdate,
	) error
	FindAddressById(
		ctx context.Context,
		id, userId int,
	) (*usermodel.UserAddress, error)
}

type updateAddressBusiness struct {
	store UpdateAddressStore
}

func NewUpdateAddressBusiness(store UpdateAddressStore) *updateAddressBusiness {
	return &updateAddressBusiness{store: store}
}

func (biz *updateAddressBusiness) UpdateAddress(
	ctx context.Context,
	id int,
	userId int,
	data *usermodel.UserAddressUpdate,
) error {
	if _, err := biz.store.FindAddressById(ctx, id, userId); err != nil {
		if common.IsRecordNotFound(err) {
			return common.ErrEntityNotFound("User Address", err)
		}
		return common.ErrorDB(err)
	}
	if err := biz.store.UpdateAddress(ctx, userId, id, data); err != nil {
		return common.ErrCannotUpdateEntity("User Address", err)
	}
	return nil
}
