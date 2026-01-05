package sellerbusiness

import (
	"OpenMarket/common"
	sellermodel "OpenMarket/module/seller/model"
	"context"
	"errors"
)

type DeleteSellerRepo interface {
	DeleteSellerByUserId(
		ctx context.Context,
		id int,
	) error
}

type deleteSellerBusiness struct {
	repo   DeleteSellerRepo
	finder FindSellerByID
}

func NewDeleteSellerBusiness(
	repo DeleteSellerRepo,
	finder FindSellerByID,
) *deleteSellerBusiness {
	return &deleteSellerBusiness{
		repo:   repo,
		finder: finder,
	}
}

func (biz *deleteSellerBusiness) DeleteSeller(
	ctx context.Context,
	userId int,
) error {
	result, err := biz.finder.FindActiveSellerWithUserID(ctx, userId)
	if err != nil {
		return common.ErrorDB(err)
	}
	if result == nil {
		return common.ErrEntityNotFound(sellermodel.EntityName, errors.New("you don't have a shop"))
	}
	if result.Status == common.SystemStatusDeleted {
		return common.ErrInvalidState(sellermodel.EntityName, "no permission")
	}
	if err := biz.repo.DeleteSellerByUserId(ctx, userId); err != nil {
		return common.ErrCannotDeleteEntity(sellermodel.EntityName, err)
	}
	return nil
}
