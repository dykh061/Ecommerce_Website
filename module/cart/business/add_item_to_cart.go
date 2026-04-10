package cartbusiness

import (
	"OpenMarket/common"
	cartmodel "OpenMarket/module/cart/model"
	cartrepository "OpenMarket/module/cart/repository"
	"context"
	"errors"
)

type addItemToCartBusiness struct {
	txRepo cartrepository.TransactionRepo
}

func NewAddItemToCartBusiness(
	txRepo cartrepository.TransactionRepo,
) *addItemToCartBusiness {
	return &addItemToCartBusiness{
		txRepo: txRepo,
	}
}
func (biz *addItemToCartBusiness) AddItemToCart(
	ctx context.Context,
	userId, variantId, quantity int,
) error {

	if quantity <= 0 {
		return common.InvalidRequestError(
			errors.New("quantity must be greater than 0"),
		)
	}

	return biz.txRepo.WithTransaction(ctx, func(tx cartrepository.TxStore) error {

		// 1. tìm cart
		cart, err := tx.FindCart(ctx, userId)
		if err != nil {
			if !common.IsRecordNotFound(err) {
				return common.ErrorDB(err)
			}
		}

		// 2. tạo cart nếu chưa có
		if cart == nil {
			if err := tx.CreateCart(ctx, &cartmodel.CartCreate{
				UserId: userId,
			}); err != nil {
				return common.ErrorDB(err)
			}

			cart, err = tx.FindCart(ctx, userId)
			if err != nil {
				return common.ErrorDB(err)
			}
		}

		// 3. tìm item
		item, err := tx.FindCartItem(ctx, cart.Id, variantId)
		if err != nil {
			if !common.IsRecordNotFound(err) {
				return common.ErrorDB(err)
			}
		}

		// 4. update hoặc create
		if item != nil {
			return tx.UpdateCartItem(ctx, item.Id, quantity)
		}

		return tx.CreateItem(ctx, &cartmodel.CartItemCreate{
			CartId:    cart.Id,
			VariantId: variantId,
			Quantity:  quantity,
		})
	})
}
