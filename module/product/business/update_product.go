package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

type UpdateProductRepo interface {
	UpdateProduct(
		ctx context.Context,
		condition map[string]interface{},
		data *productmodel.ProductUpdate,
	) error
}

type updateProductBusiness struct {
	repo    UpdateProductRepo
	sfinder FindSellerByID
	pfinder FindProductWithIDAndSellerID
}

func NewUpdateProductBusiness(
	repo UpdateProductRepo,
	sfinder FindSellerByID,
	pfinder FindProductWithIDAndSellerID,
) *updateProductBusiness {
	return &updateProductBusiness{
		repo:    repo,
		sfinder: sfinder,
		pfinder: pfinder,
	}
}
func (biz *updateProductBusiness) UpdateProduct(
	ctx context.Context,
	userId int,
	productId int,
	data *productmodel.ProductUpdate,
) error {
	seller, err := biz.sfinder.FindActiveSellerWithUserID(ctx, userId)
	if err != nil {
		return common.ErrCannotUpdateEntity(productmodel.EntityName, err)
	}
	if _, err := biz.pfinder.FindProductByIdWithSellerID(ctx, productId, seller.Id); err != nil {
		return common.ErrCannotUpdateEntity(productmodel.EntityName, err)
	}
	if err := biz.repo.UpdateProduct(ctx, map[string]interface{}{
		"id":     productId,
		"status": common.SystemStatusActive,
	}, data); err != nil {
		return common.ErrCannotUpdateEntity(productmodel.EntityName, err)
	}
	return nil
}
