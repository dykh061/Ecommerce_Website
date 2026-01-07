package productbusiness

import (
	productmodel "OpenMarket/module/product/model"
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type FindSellerByID interface {
	FindActiveSellerWithUserID(
		ctx context.Context,
		userid int,
	) (*sellermodel.Seller, error)
}

type FindProductWithIDAndSellerID interface {
	FindProductByIdWithSellerID(
		ctx context.Context,
		id int,
		sellerID int,
	) (*productmodel.Product, error)
}

type GetProduct interface {
	GetProductById(
		ctx context.Context,
		productId int,
	) (*productmodel.Product, error)
}

type GetImagesRepo interface {
	GetImages(
		ctx context.Context,
		productID int,
	) ([]string, error)
}
