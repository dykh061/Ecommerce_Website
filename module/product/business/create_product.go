package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type CreateProductRepo interface {
	CreateProduct(
		ctx context.Context,
		data *productmodel.ProductCreate,
	) error
}

type FindSellerRepo interface {
	FindActiveSellerWithUserID(
		ctx context.Context,
		userid int,
	) (*sellermodel.Seller, error)
}

type createProductBusiness struct {
	productRepo CreateProductRepo
	sellerRepo  FindSellerRepo
}

func NewCreateProductBusiness(
	productRepo CreateProductRepo,
	sellerRepo FindSellerRepo,
) *createProductBusiness {
	return &createProductBusiness{
		productRepo: productRepo,
		sellerRepo:  sellerRepo,
	}
}

func (biz *createProductBusiness) CreateProduct(
	ctx context.Context,
	userID int,
	data *productmodel.ProductCreate,
) error {

	seller, err := biz.sellerRepo.FindActiveSellerWithUserID(ctx, userID)
	if err != nil || seller == nil {
		return common.ErrEntityNotFound("Shop", err)
	}

	data.SellerID = seller.Id

	if err := biz.productRepo.CreateProduct(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("Product", err)
	}

	return nil
}
