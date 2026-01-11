package orderstorage

import (
	cartrepository "OpenMarket/module/cart/repository"
	ordermodel "OpenMarket/module/order/model"
	productmodel "OpenMarket/module/product/model"
	productrepository "OpenMarket/module/product/repository"

	"context"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type txStore struct {
	db            *gorm.DB
	orderStore    *sqlStore
	stockStore    productrepository.AdjustStockStorage
	variantReader productrepository.VariantReader
	cartStore     cartrepository.CartCleaner
}

func (t *txStore) CreateOrder(
	ctx context.Context,
	data *ordermodel.Order,
) error {
	return t.orderStore.CreateOrder(ctx, data)
}

func (t *txStore) CreateOrderItem(
	ctx context.Context,
	item *ordermodel.OrderItemCreate,
) error {
	return t.orderStore.CreateOrderItem(ctx, item)
}

func (t *txStore) AdjustVariantStock(
	ctx context.Context,
	variantID int,
	by int,
) error {
	return t.stockStore.AdjustVariantStock(ctx, variantID, by)
}

func (t *txStore) MarkOrderAsPaid(
	ctx context.Context,
	orderID int,
) error {
	return t.orderStore.UpdateOrderStatus(
		ctx,
		orderID,
		ordermodel.OrderConfirmed,
	)
}

func (t *txStore) DeleteCart(
	ctx context.Context,
	userId int,
) error {
	return t.cartStore.DeleteCart(ctx, userId)
}

func (t *txStore) FindVariantByID(
	ctx context.Context,
	id int,
) (*productmodel.Variant, error) {
	return t.variantReader.FindVariantByID(ctx, id)
}

func (t *txStore) UpTotalAmount(
	ctx context.Context,
	totalAmount decimal.Decimal,
	id int,
) error {
	return t.orderStore.UpTotalAmount(ctx, totalAmount, id)
}
