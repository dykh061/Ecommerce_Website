package sellerbusiness

import (
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type FindSellerByID interface {
	FindActiveSellerWithUserID(
		ctx context.Context,
		userid int,
	) (*sellermodel.Seller, error)
}
