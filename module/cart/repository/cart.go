package cartrepository

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

type CartCleaner interface {
	DeleteCart(
		ctx context.Context,
		userId int,
	) error
	FindCart(
		ctx context.Context,
		userId int,
	) (*cartmodel.Cart, error)
	ListCartItems(
		ctx context.Context,
		cartId int,
	) ([]cartmodel.CartItem, error)
}
