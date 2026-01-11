package orderbusiness

import (
	"OpenMarket/common"
	ordermodel "OpenMarket/module/order/model"
	orderrepository "OpenMarket/module/order/repository"
	"context"
	"errors"

	"github.com/shopspring/decimal"
)

type createOrderBusiness struct {
	txRepo orderrepository.TransactionRepo
}

func NewCreateOrderBusiness(txRepo orderrepository.TransactionRepo) *createOrderBusiness {
	return &createOrderBusiness{txRepo: txRepo}
}

func (biz *createOrderBusiness) CreateOrder(
	ctx context.Context,
	userID int,
	addressID int,
) error {

	return biz.txRepo.WithTransaction(ctx, func(tx orderrepository.TxStore) error {

		// 1️⃣ Find cart
		cart, err := tx.FindCart(ctx, userID)
		if err != nil {
			return err
		}

		// 2️⃣ List cart items
		items, err := tx.ListCartItems(ctx, cart.Id)
		if err != nil {
			return err
		}

		if len(items) == 0 {
			return common.InvalidRequestError(errors.New("cart is empty"))
		}

		// 3️⃣ Validate address
		address, err := tx.FindAddressById(ctx, addressID, userID)
		if err != nil {
			return err
		}

		user, err := tx.FindDataWithCondition(ctx, map[string]interface{}{
			"id": userID,
		})
		if err != nil {
			return err
		}

		// 4️⃣ Create order
		order := &ordermodel.Order{
			UserId:      userID,
			TotalAmount: decimal.Zero,
			Status:      ordermodel.OrderPending,
		}
		if err := tx.CreateOrder(ctx, order); err != nil {
			return err
		}

		// 5️⃣ Snapshot address
		if err := tx.CreateAddress(ctx, &ordermodel.OrderAddressCreate{
			OrderId:  order.Id,
			Address:  address.Address,
			City:     address.City,
			Phone:    user.Phone,
			FullName: user.Name,
		}); err != nil {
			return err
		}

		total := decimal.Zero

		// 6️⃣ Create order_items từ cart_items
		for _, ci := range items {

			if ci.Quantity <= 0 {
				return common.InvalidRequestError(errors.New("invalid quantity"))
			}

			variant, err := tx.FindVariantByID(ctx, ci.VariantId)
			if err != nil {
				return err
			}

			sub := variant.Price.Mul(decimal.NewFromInt(int64(ci.Quantity)))
			total = total.Add(sub)

			if err := tx.CreateOrderItem(ctx, &ordermodel.OrderItemCreate{
				OrderId:   order.Id,
				VariantId: ci.VariantId,
				Quantity:  ci.Quantity,
				Price:     variant.Price,
			}); err != nil {
				return err
			}

			if err := tx.AdjustVariantStock(ctx, ci.VariantId, -ci.Quantity); err != nil {
				return err
			}
		}

		// 7️⃣ Update total
		if err := tx.UpTotalAmount(ctx, total, order.Id); err != nil {
			return err
		}

		// 8️⃣ Clear cart
		return tx.DeleteCart(ctx, userID)
	})
}
