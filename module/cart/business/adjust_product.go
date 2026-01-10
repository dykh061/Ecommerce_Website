package cartbusiness

import (
	"OpenMarket/common"
	"context"
)

type adjustProductInCartBusiness struct {
	repo     AdjustCartItemRepo
	findItem FindCartItem
	finder   GetCart
}

func NewAdjustProductInCartBusiness(
	repo AdjustCartItemRepo,
	finder GetCart,
	findItem FindCartItem,
) *adjustProductInCartBusiness {
	return &adjustProductInCartBusiness{
		repo:     repo,
		finder:   finder,
		findItem: findItem,
	}
}

func (biz *adjustProductInCartBusiness) AdjustProductInCart(
	ctx context.Context,
	quantity, userId, variantId int,
) error {
	cart, err := biz.finder.FindCartWithId(ctx, userId)
	if err != nil {
		return common.ErrCannotListEntity("cart item", err)
	}
	cartItem, err := biz.findItem.FindCartItemWithId(ctx, cart.Id, variantId)
	if err != nil {
		return common.ErrCannotListEntity("cart item", err)
	}
	if cartItem == nil {
		return common.ErrCannotListEntity("cart item", err)
	}
	return biz.repo.AdjustCartItem(ctx, cartItem.Id, quantity)

}
