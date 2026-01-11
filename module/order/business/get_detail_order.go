package orderbusiness

import (
	ordermodel "OpenMarket/module/order/model"

	"context"
)

type GetDetailOrderStore interface {
	ListOrderItemsByOrderID(
		ctx context.Context,
		orderID int,
	) ([]ordermodel.OrderItem, error)
	FindOrderByID(
		ctx context.Context,
		id int,
	) (*ordermodel.Order, error)
}

type getDetailOrderBusiness struct {
	store GetDetailOrderStore
}

func NewGetDetailOrderBusiness(store GetDetailOrderStore) *getDetailOrderBusiness {
	return &getDetailOrderBusiness{store: store}
}

func (biz *getDetailOrderBusiness) GetDetailOrder(
	ctx context.Context,
	orderId int,
) (*ordermodel.OrderDetail, error) {
	order, err := biz.store.FindOrderByID(ctx, orderId)
	if err != nil {
		return nil, err
	}
	items, err := biz.store.ListOrderItemsByOrderID(ctx, orderId)
	if err != nil {
		return nil, err
	}

	orderDetail := ordermodel.OrderDetail{
		Id:          orderId,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}
	for _, it := range items {
		orderItem := ordermodel.OrderWithItems{
			VariantId: it.VariantId,
			Quantity:  it.Quantity,
			Price:     it.Price,
		}
		orderDetail.Items = append(orderDetail.Items, orderItem)
	}
	return &orderDetail, nil
}
