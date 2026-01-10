package cartrepository

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

type TxStore interface {
	CreateCart(
		ctx context.Context,
		data *cartmodel.CartCreate,
	) error

	CreateItem(
		ctx context.Context,
		item *cartmodel.CartItemCreate,
	) error
	UpdateCartItem(
		ctx context.Context,
		id, by int,
	) error
	FindCart(
		ctx context.Context,
		userId int,
	) (*cartmodel.Cart, error)
	FindCartItem(
		ctx context.Context,
		cartId, variantId int,
	) (*cartmodel.CartItem, error)
	DeleteCartItem(ctx context.Context, id int) error
}

type TransactionRepo interface {
	WithTransaction(
		ctx context.Context,
		fn func(tx TxStore) error,
	) error
}
