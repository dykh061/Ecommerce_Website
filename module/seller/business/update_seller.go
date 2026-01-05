package sellerbusiness

import (
	"OpenMarket/common"
	sellermodel "OpenMarket/module/seller/model"
	"context"
	"errors"
)

type UpdateSellerRepo interface {
	UpdateSeller(
		ctx context.Context,
		userId int,
		data sellermodel.SellerUpdate,
	) error
}

type updateSellerBusiness struct {
	finder FindSellerByID
	repo   UpdateSellerRepo
}

func NewUpdateSellerBusiness(
	finder FindSellerByID,
	repo UpdateSellerRepo,
) *updateSellerBusiness {
	return &updateSellerBusiness{
		finder: finder,
		repo:   repo,
	}
}

func (biz *updateSellerBusiness) UpdateSeller(
	cxt context.Context,
	userId int,
	data sellermodel.SellerUpdate,
) error {
	oldData, err := biz.finder.FindActiveSellerWithUserID(cxt, userId)
	if err != nil {
		return common.ErrorDB(err)
	}
	if oldData == nil {
		return common.ErrEntityNotFound(sellermodel.EntityName, errors.New("seller not found"))
	}
	if oldData.Status == common.SystemStatusDeleted {
		return common.ErrInvalidState(sellermodel.EntityName, "no permission")
	}
	if err := biz.repo.UpdateSeller(cxt, userId, data); err != nil {
		return common.ErrCannotUpdateEntity(sellermodel.EntityName, err)
	}
	return nil

}
