package orderrepository

import (
	ordermodel "OpenMarket/module/order/model"
	productmodel "OpenMarket/module/product/model"
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
}

type TransactionRepo interface {
	WithTransaction(
		ctx context.Context,
		fn func(tx TxStore) error,
	) error
}
