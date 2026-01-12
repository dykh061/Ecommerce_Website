package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
	"errors"
)

type GetVariantWithAttrsRepo interface {
	GetVariantWithAttributes(
		ctx context.Context,
		variantID int,
		productID int,
	) (*productmodel.VariantDetailFull, error)
}

type getSellerVariantDetailBusiness struct {
	sfinder     FindSellerByID
	pfinder     FindProductWithIDAndSellerID
	variantRepo GetVariantWithAttrsRepo
}

func NewGetSellerVariantDetailBusiness(
	sfinder FindSellerByID,
	pfinder FindProductWithIDAndSellerID,
	variantRepo GetVariantWithAttrsRepo,
) *getSellerVariantDetailBusiness {
	return &getSellerVariantDetailBusiness{
		sfinder:     sfinder,
		pfinder:     pfinder,
		variantRepo: variantRepo,
	}
}

func (biz *getSellerVariantDetailBusiness) GetSellerVariantDetail(
	ctx context.Context,
	userID int,
	productID int,
	variantID int,
) (*productmodel.VariantDetailFull, error) {
	// 1. Find seller by userID
	seller, err := biz.sfinder.FindActiveSellerWithUserID(ctx, userID)
	if err != nil {
		return nil, common.ErrForbidden(errors.New("user is not a seller"))
	}

	// 2. Find product with seller ownership check
	product, err := biz.pfinder.FindProductByIdWithSellerID(ctx, productID, seller.Id)
	if err != nil {
		return nil, common.ErrForbidden(errors.New("product does not belong to this seller"))
	}
	if product == nil {
		return nil, common.ErrEntityNotFound(productmodel.EntityName, errors.New("product not found"))
	}

	// 3. Get variant with attributes
	variant, err := biz.variantRepo.GetVariantWithAttributes(ctx, variantID, productID)
	if err != nil {
		return nil, common.ErrEntityNotFound("Variant", err)
	}
	if variant == nil {
		return nil, common.ErrEntityNotFound("Variant", errors.New("variant not found"))
	}

	return variant, nil
}
