package sellerbusiness

import (
	"OpenMarket/common"
	sellermodel "OpenMarket/module/seller/model"
	"context"
	"errors"
)

type RegisterSellerRepo interface {
	RegisterSeller(
		ctx context.Context,
		data *sellermodel.SellerCreate,
	) error
}
type FindSellerRepo interface {
	FindSeller(
		ctx context.Context,
		userid int,
	) (*sellermodel.Seller, error)
}

type createSellerStore struct {
	crepo RegisterSellerRepo
	frepo FindSellerRepo
}

func NewCreateSellerBusiness(crepo RegisterSellerRepo, frepo FindSellerRepo) *createSellerStore {
	return &createSellerStore{crepo: crepo, frepo: frepo}
}

func (biz *createSellerStore) CreateSeller(ctx context.Context, userId int, data *sellermodel.SellerCreate) error {
	result, err := biz.frepo.FindSeller(ctx, userId)
	if err != nil {
		if !common.IsRecordNotFound(err) {
			return common.ErrorDB(err)
		}
	}
	if result != nil {
		if result.Status == common.SystemStatusActive {
			return common.ErrUserAlreadyHasSeller(errors.New("user already has a seller"))
		}
		if result.Status == common.SystemStatusDeleted {
			return common.ErrSellerWasSoftDeleted(errors.New("you are not allowed to create a shop"))
		}
	}
	if err := biz.crepo.RegisterSeller(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(sellermodel.EntityName, err)
	}
	return nil
}
