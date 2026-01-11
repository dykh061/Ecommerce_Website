package orderbusiness

import (
	"OpenMarket/common"
	ordermodel "OpenMarket/module/order/model"
	"context"
	"errors"
	"time"
)

type CancelOrderStore interface {
	FindOrderByID(
		ctx context.Context,
		id int,
	) (*ordermodel.Order, error)
	CancelOrder(
		ctx context.Context,
		id int,
		reason string,
	) error
}

type cancelOrderBusiness struct {
	store CancelOrderStore
}

func NewCancelOrderBusiness(store CancelOrderStore) *cancelOrderBusiness {
	return &cancelOrderBusiness{store: store}
}
func (business *cancelOrderBusiness) CancelOrder(
	ctx context.Context,
	id int,
	reason string,
) error {
	order, err := business.store.FindOrderByID(ctx, id)
	if err != nil {
		return err
	}
	if order.Status != ordermodel.OrderPending {
		return common.InvalidRequestError(
			errors.New("order cannot be cancelled"),
		)
	}
	if order.CreatedAt == nil {
		return common.ErrInternal(errors.New("created_at is nil"))
	}
	// từ thời điểm tạo đơn since hiện tại mà quá 24h thì không được hủy đơn
	if time.Since(*order.CreatedAt) > 24*time.Hour {
		return common.InvalidRequestError(
			errors.New("order can only be cancelled within 24 hours"),
		)
	}
	if err := business.store.CancelOrder(ctx, id, reason); err != nil {
		return err
	}
	return nil
}
