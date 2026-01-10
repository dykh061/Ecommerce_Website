package cartbusiness

import (
	cartmodel "OpenMarket/module/cart/model"
	"context"
)

type findCartBusiness struct {
	repo FindCartRepo
}

func NewFindCartBusiness(repo FindCartRepo) *findCartBusiness {
	return &findCartBusiness{repo: repo}
}

func (biz *findCartBusiness) FindCart(
	ctx context.Context,
	userId int,
) (*cartmodel.Cart, error) {
	return biz.repo.FindCartWithId(ctx, userId)
}
