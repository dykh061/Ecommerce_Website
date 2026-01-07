package productbusiness

import (
	"OpenMarket/common"
	"context"
)

type AdjustStockRepo interface {
	AdjustStock(
		ctx context.Context,
		variantId int,
		by int,
	) error
}

type adjustStockBusiness struct {
	repo AdjustStockRepo
}

func NewAdjustStockBusiness(
	repo AdjustStockRepo,
) *adjustStockBusiness {
	return &adjustStockBusiness{
		repo: repo,
	}
}

func (biz *adjustStockBusiness) AdjustStock(
	ctx context.Context,
	variantId int,
	by int,
) error {
	if by == 0 {
		return nil
	}
	if err := biz.repo.AdjustStock(ctx, variantId, by); err != nil {
		return common.ErrCannotUpdateEntity("Variant", err)
	}
	return nil
}
