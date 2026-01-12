package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

type GetProductAttributesRepo interface {
	GetProductAttributes(
		ctx context.Context,
		productID int,
	) ([]productmodel.ProductAttribute, error)
}

type getProductAttributesBusiness struct {
	productFinder GetProduct
	attrRepo      GetProductAttributesRepo
}

func NewGetProductAttributesBusiness(
	productFinder GetProduct,
	attrRepo GetProductAttributesRepo,
) *getProductAttributesBusiness {
	return &getProductAttributesBusiness{
		productFinder: productFinder,
		attrRepo:      attrRepo,
	}
}

func (biz *getProductAttributesBusiness) GetProductAttributes(
	ctx context.Context,
	productID int,
) ([]productmodel.ProductAttribute, error) {
	// 1. Check product exists
	product, err := biz.productFinder.GetProductById(ctx, productID)
	if err != nil || product == nil {
		return nil, common.ErrEntityNotFound(productmodel.EntityName, err)
	}

	// 2. Get attributes
	attributes, err := biz.attrRepo.GetProductAttributes(ctx, productID)
	if err != nil {
		return nil, common.ErrCannotListEntity("Attribute", err)
	}

	return attributes, nil
}
