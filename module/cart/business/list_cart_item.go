package cartbusiness

import (
	"OpenMarket/common"
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

type ListCartItemRepo interface {
	FindCart(
		ctx context.Context,
		userId int,
	) (*cartmodel.Cart, error)

	ListCartItems(
		ctx context.Context,
		cartId int,
	) ([]cartmodel.CartItem, error)
}

type listCartItemBusiness struct {
	repo ListCartItemRepo
}

func NewListCartItemBusiness(repo ListCartItemRepo) *listCartItemBusiness {
	return &listCartItemBusiness{repo: repo}
}

func (biz *listCartItemBusiness) ListCartItem(
	ctx context.Context,
	userId int,
) (*cartmodel.CartView, error) {
	cart, err := biz.repo.FindCart(ctx, userId)
	if err != nil {
		if common.IsRecordNotFound(err) {
			return &cartmodel.CartView{
				Items: []cartmodel.CartItemView{},
			}, nil
		}
		return nil, err
	}
	items, err := biz.repo.ListCartItems(ctx, cart.Id)
	if err != nil {
		return nil, err
	}
	result := &cartmodel.CartView{
		Id: cart.Id,
	}
	for _, item := range items {
		result.Items = append(result.Items, cartmodel.CartItemView{
			VariantId: item.VariantId,
			Quantity:  item.Quantity,
		})
	}
	return result, nil
}
