package cartbusiness

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CreateCartRepo interface {
	MakeCart(
		ctx context.Context,
		userId int,
	) error
}
type FindCartRepo interface {
	FindCartWithId(
		ctx context.Context,
		userId int,
	) (*cartmodel.Cart, error)
}

type createCartBusiness struct {
	repo CreateCartRepo
	find FindCartRepo
}

func NewCreateCartBusiness(repo CreateCartRepo, find FindCartRepo) *createCartBusiness {
	return &createCartBusiness{repo: repo, find: find}
}
func (biz *createCartBusiness) CreateCart(
	ctx context.Context,
	userId int,
) error {
	data, err := biz.find.FindCartWithId(ctx, userId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if data != nil {
		return nil
	}
	if err := biz.repo.MakeCart(ctx, userId); err != nil {
		return err
	}
	return nil
}
