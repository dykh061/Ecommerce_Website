package orderbusiness

import (
	ordermodel "OpenMarket/module/order/model"

	"context"
)

type GetDetailOrderStore interface {
	ListOrderItemsByOrderID(
		ctx context.Context,
		orderID int,
	) ([]ordermodel.OrderItemDetailRow, error)
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
	itemByID := make(map[int]*ordermodel.OrderWithItems)
	itemOrder := make([]int, 0)
	for _, row := range items {
		it, ok := itemByID[row.OrderItemID]
		if !ok {
			name := ""
			if row.ProductName != nil {
				name = *row.ProductName
			}
			it = &ordermodel.OrderWithItems{
				VariantId:  row.VariantId,
				Quantity:   row.Quantity,
				Price:      row.Price,
				Name:       name,
				Image:      row.ImageURL,
				Attributes: map[string]string{},
			}
			itemByID[row.OrderItemID] = it
			itemOrder = append(itemOrder, row.OrderItemID)
		}
		if row.AttributeName != nil && row.AttributeValue != nil {
			it.Attributes[*row.AttributeName] = *row.AttributeValue
		}
	}
	for _, id := range itemOrder {
		orderDetail.Items = append(orderDetail.Items, *itemByID[id])
	}
	return &orderDetail, nil
}
