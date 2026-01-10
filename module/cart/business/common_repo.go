package cartbusiness

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

type GetCart interface {
	FindCartWithId(
		ctx context.Context,
		userId int,
	) (*cartmodel.Cart, error)
}

type CreateCartItemRepo interface {
	CreateItem(
		ctx context.Context,
		item *cartmodel.CartItemCreate,
	) error
}
type CreateCart interface {
	CreateCart(
		ctx context.Context,
		userId int,
	) error
}
type FindCartItem interface {
	FindCartItemWithId(
		ctx context.Context,
		cartId, variantId int,
	) (*cartmodel.CartItem, error)
}

type AdjustCartItemRepo interface {
	AdjustCartItem(
		ctx context.Context,
		id, quantity int,
	) error
}
