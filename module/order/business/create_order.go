package orderbusiness

import (
	"OpenMarket/common"
	ordermodel "OpenMarket/module/order/model"
	orderrepository "OpenMarket/module/order/repository"
	"context"
	"errors"

	"github.com/shopspring/decimal"
)

type CreateOrderItem struct {
	VariantID int
	Quantity  int
}

type createOrderBusiness struct {
	txRepo orderrepository.TransactionRepo
}

func NewCreateOrderBusiness(txRepo orderrepository.TransactionRepo) *createOrderBusiness {
	return &createOrderBusiness{txRepo: txRepo}
}

func (biz *createOrderBusiness) CreateOrder(
	ctx context.Context,
	userID int,
	items []CreateOrderItem,
) error {
	if len(items) == 0 {
		return common.InvalidRequestError(errors.New("order items is empty"))
	}
	return biz.txRepo.WithTransaction(ctx, func(tx orderrepository.TxStore) error {

		total := decimal.Zero

		//  Tạo order trước (total tạm = 0)
		order := &ordermodel.Order{
			UserId:      userID,
			TotalAmount: decimal.Zero,
			Status:      ordermodel.OrderPending,
		}

		if err := tx.CreateOrder(ctx, order); err != nil {
			return err
		}

		//  Với mỗi item → lấy variant + tính tiền + tạo order_item
		for _, item := range items {

			if item.Quantity <= 0 {
				return common.InvalidRequestError(errors.New("invalid quantity"))
			}

			variant, err := tx.FindVariantByID(
				ctx,
				item.VariantID,
			)
			if err != nil {
				return err
			}

			price := variant.Price
			subTotal := price.Mul(decimal.NewFromInt(int64(item.Quantity)))
			total = total.Add(subTotal)

			//  Tạo order_item (price = DB price)
			if err := tx.CreateOrderItem(ctx, &ordermodel.OrderItemCreate{
				OrderId:   order.Id,
				VariantId: variant.Id,
				Quantity:  item.Quantity,
				Price:     price,
			}); err != nil {
				return err
			}

			if err := tx.AdjustVariantStock(
				ctx,
				variant.Id,
				-item.Quantity,
			); err != nil {
				return err
			}
		}

		//  Update lại total_amount cho order
		if err := tx.UpTotalAmount(ctx, total, order.Id); err != nil {
			return err
		}

		return nil
	})
}
