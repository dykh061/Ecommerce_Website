package orderstorage

import (
	cartmodel "OpenMarket/module/cart/model"
	cartrepository "OpenMarket/module/cart/repository"
	ordermodel "OpenMarket/module/order/model"
	productmodel "OpenMarket/module/product/model"
	productrepository "OpenMarket/module/product/repository"
	usermodel "OpenMarket/module/user/model"
	srrepository "OpenMarket/module/user/repository"

	"context"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type txStore struct {
	db            *gorm.DB
	orderStore    *sqlStore
	userStore     srrepository.UserStore
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

func (t *txStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {
	return t.userStore.FindDataWithCondition(context, condition, moreKeys...)
}

func (t *txStore) FindAddressById(
	ctx context.Context,
	id, userId int,
) (*usermodel.UserAddress, error) {
	return t.userStore.FindAddressById(ctx, id, userId)
}

func (t *txStore) CreateAddress(
	ctx context.Context,
	data *ordermodel.OrderAddressCreate,
) error {
	return t.orderStore.CreateAddress(ctx, data)
}

func (t *txStore) FindCart(ctx context.Context, userId int) (*cartmodel.Cart, error) {
	return t.cartStore.FindCart(ctx, userId)
}

func (t *txStore) ListCartItems(
	ctx context.Context,
	cartId int,
) ([]cartmodel.CartItem, error) {
	return t.cartStore.ListCartItems(ctx, cartId)
}
