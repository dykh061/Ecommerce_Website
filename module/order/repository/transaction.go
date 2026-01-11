package orderrepository

import (
	cartmodel "OpenMarket/module/cart/model"
	ordermodel "OpenMarket/module/order/model"
	productmodel "OpenMarket/module/product/model"
	usermodel "OpenMarket/module/user/model"
	"context"

	"github.com/shopspring/decimal"
)

type TxStore interface {
	CreateOrderItem(
		ctx context.Context,
		item *ordermodel.OrderItemCreate,
	) error

	CreateOrder(
		ctx context.Context,
		data *ordermodel.Order,
	) error

	AdjustVariantStock(
		ctx context.Context,
		variantID int,
		by int,
	) error

	FindVariantByID(
		ctx context.Context,
		id int,
	) (*productmodel.Variant, error)
	// order
	MarkOrderAsPaid(ctx context.Context, orderID int) error

	// cart
	DeleteCart(
		ctx context.Context,
		userId int,
	) error
	UpTotalAmount(
		ctx context.Context,
		totalAmount decimal.Decimal,
		id int,
	) error
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
	FindAddressById(
		ctx context.Context,
		id, userId int,
	) (*usermodel.UserAddress, error)
	CreateAddress(
		ctx context.Context,
		data *ordermodel.OrderAddressCreate,
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

type TransactionRepo interface {
	WithTransaction(
		ctx context.Context,
		fn func(tx TxStore) error,
	) error
}
