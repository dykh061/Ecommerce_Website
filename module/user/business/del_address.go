package userbusiness

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type DeleteAddressStore interface {
	FindAddressById(
		ctx context.Context,
		id, userId int,
	) (*usermodel.UserAddress, error)
	DeleteAddress(
		ctx context.Context,
		id, userId int,
	) error
}
type deleteAddressBusiness struct {
	store DeleteAddressStore
}

func NewDeleteAddressBusiness(store DeleteAddressStore) *deleteAddressBusiness {
	return &deleteAddressBusiness{store: store}
}

func (biz *deleteAddressBusiness) DeleteAddress(
	ctx context.Context,
	id int,
	userId int,
) error {
	if _, err := biz.store.FindAddressById(ctx, id, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.ErrEntityNotFound("User Address", err)
		}
		return common.ErrorDB(err)
	}
	if err := biz.store.DeleteAddress(ctx, id, userId); err != nil {
		return common.ErrCannotDeleteEntity("User Address", err)
	}
	return nil
}
