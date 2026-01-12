package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

type CreateVariantRepo interface {
	CreateVariantWithAttributes(
		ctx context.Context,
		productID int,
		data *productmodel.VariantCreate,
	) (*productmodel.VariantCreate, error)
}

type createVariantBusiness struct {
	variantRepo   CreateVariantRepo
	sellerFinder  FindSellerByID
	productFinder FindProductWithIDAndSellerID
}

func NewCreateVariantBusiness(
	variantRepo CreateVariantRepo,
	sellerFinder FindSellerByID,
	productFinder FindProductWithIDAndSellerID,
) *createVariantBusiness {
	return &createVariantBusiness{
		variantRepo:   variantRepo,
		sellerFinder:  sellerFinder,
		productFinder: productFinder,
	}
}

func (biz *createVariantBusiness) CreateVariant(
	ctx context.Context,
	userID int,
	productID int,
	data *productmodel.VariantCreate,
) (*productmodel.VariantCreate, error) {
	seller, err := biz.sellerFinder.FindActiveSellerWithUserID(ctx, userID)
	if err != nil || seller == nil {
		return nil, common.ErrEntityNotFound("Shop", err)
	}
	product, err := biz.productFinder.FindProductByIdWithSellerID(ctx, productID, seller.Id)
	if err != nil || product == nil {
		return nil, common.ErrEntityNotFound("Product", err)
	}
	if err := data.Validate(); err != nil {
		return nil, common.InvalidRequestError(err)
	}
	data.Sku = BuildSKUWithUID(seller.Id, productID)
	variant, err := biz.variantRepo.CreateVariantWithAttributes(ctx, productID, data)
	if err != nil {
		return nil, common.ErrCannotCreateEntity("Variant", err)
	}
	return variant, nil
}
