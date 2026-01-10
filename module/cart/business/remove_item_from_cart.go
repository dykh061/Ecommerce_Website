package cartbusiness

import (
	"OpenMarket/common"
	cartrepository "OpenMarket/module/cart/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

type removeItemFromCartBusiness struct {
	txRepo cartrepository.TransactionRepo
}

func NewRemoveItemFromCartBusiness(
	txRepo cartrepository.TransactionRepo,
) *removeItemFromCartBusiness {
	return &removeItemFromCartBusiness{txRepo: txRepo}
}

func (biz *removeItemFromCartBusiness) RemoveItemFromCart(
	ctx context.Context,
	userId, variantId int,
) error {
	return biz.txRepo.WithTransaction(ctx, func(tx cartrepository.TxStore) error {
		cart, err := tx.FindCart(ctx, userId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return common.InvalidRequestError(errors.New("cart not found"))
			}
			return common.ErrorDB(err)
		}

		item, err := tx.FindCartItem(ctx, cart.Id, variantId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return common.InvalidRequestError(errors.New("item not found in cart"))
			}
			return common.ErrorDB(err)
		}

		if err := tx.DeleteCartItem(ctx, item.Id); err != nil {
			return common.ErrorDB(err)
		}
		return nil
	})
}
