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
	) (*productmodel.ProductCreate, error)
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
) (*productmodel.ProductCreate, error) {

	seller, err := biz.sellerRepo.FindActiveSellerWithUserID(ctx, userID)
	if err != nil || seller == nil {
		return nil, common.ErrEntityNotFound("Shop", err)
	}

	data.SellerID = seller.Id

	product, err := biz.productRepo.CreateProduct(ctx, data)
	if err != nil {
		return nil, common.ErrCannotCreateEntity("Product", err)
	}

	return product, nil
}
