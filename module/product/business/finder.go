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
