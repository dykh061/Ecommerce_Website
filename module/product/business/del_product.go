package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

type DeleteProductRepo interface {
	DeleteProduct(
		ctx context.Context,
		productId int,
	) error
}

type deleteProductBusiness struct {
	storage DeleteProductRepo
	sfinder FindSellerByID
	pfinder FindProductWithIDAndSellerID
}

func NewDeleteProductBusiness(
	storage DeleteProductRepo,
	sfinder FindSellerByID,
	pfinder FindProductWithIDAndSellerID,
) *deleteProductBusiness {
	return &deleteProductBusiness{
		storage: storage,
		sfinder: sfinder,
		pfinder: pfinder,
	}
}

func (biz *deleteProductBusiness) DeleteProduct(
	ctx context.Context,
	userId int,
	productId int,
) error {
	seller, err := biz.sfinder.FindActiveSellerWithUserID(ctx, userId)
	if err != nil {
		return common.ErrCannotDeleteEntity(productmodel.EntityName, err)
	}
	if _, err := biz.pfinder.FindProductByIdWithSellerID(ctx, productId, seller.Id); err != nil {
		return common.ErrCannotDeleteEntity(productmodel.EntityName, err)
	}
	if err := biz.storage.DeleteProduct(ctx, productId); err != nil {
		return common.ErrCannotDeleteEntity(productmodel.EntityName, err)
	}
	return nil
}
