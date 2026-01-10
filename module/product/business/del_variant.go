package productbusiness

import (
	"OpenMarket/common"
	"context"
)

type DeleteVariantRepo interface {
	DeleteVariant(
		ctx context.Context,
		variantId int,
	) error
}

type deleteVariantRepo struct {
	repo    DeleteVariantRepo
	sfinder FindSellerByID
	pfinder FindProductWithIDAndSellerID
}

func NewDeleteVariantBusiness(
	repo DeleteVariantRepo,
	sfinder FindSellerByID,
	pfinder FindProductWithIDAndSellerID,
) *deleteVariantRepo {
	return &deleteVariantRepo{
		repo:    repo,
		sfinder: sfinder,
		pfinder: pfinder,
	}
}

func (biz *deleteVariantRepo) DeleteVariant(
	ctx context.Context,
	userId int,
	productId int,
	variantId int,
) error {
	seller, err := biz.sfinder.FindActiveSellerWithUserID(ctx, userId)
	if err != nil {
		return common.ErrCannotDeleteEntity("variant", err)
	}
	if _, err := biz.pfinder.FindProductByIdWithSellerID(ctx, productId, seller.Id); err != nil {
		return common.ErrCannotDeleteEntity("variant", err)
	}
	if err := biz.repo.DeleteVariant(ctx, variantId); err != nil {
		return common.ErrCannotDeleteEntity("variant", err)
	}
	return nil
}
