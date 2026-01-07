package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

type UpdateVariantRepo interface {
	UpdateVariant(
		ctx context.Context,
		variantId int,
		condition map[string]interface{},
		data *productmodel.VariantUpdate,
	) error
}

type updateVariantBusiness struct {
	repo    UpdateVariantRepo
	sfinder FindSellerByID
	vfinder FindProductWithIDAndSellerID
}

func NewUpdateVariantBusiness(
	repo UpdateVariantRepo,
	sfinder FindSellerByID,
	vfinder FindProductWithIDAndSellerID,
) *updateVariantBusiness {
	return &updateVariantBusiness{
		repo:    repo,
		sfinder: sfinder,
		vfinder: vfinder,
	}
}

func (biz *updateVariantBusiness) UpdateVariant(
	ctx context.Context,
	userId int,
	productId int,
	variantId int,
	data *productmodel.VariantUpdate,
) error {
	seller, err := biz.sfinder.FindActiveSellerWithUserID(ctx, userId)
	if err != nil {
		return common.ErrCannotUpdateEntity("variant", err)
	}
	if _, err := biz.vfinder.FindProductByIdWithSellerID(ctx, productId, seller.Id); err != nil {
		return common.ErrCannotUpdateEntity("variant", err)
	}
	if err := biz.repo.UpdateVariant(ctx, variantId, map[string]interface{}{
		"id":         variantId,
		"product_id": productId,
		"status":     common.SystemStatusActive,
	}, data); err != nil {
		return common.ErrCannotUpdateEntity("variant", err)
	}
	return nil
}
