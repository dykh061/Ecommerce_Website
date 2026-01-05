package sellerbusiness

import (
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type GetSellerRepo interface {
	GetSellerWithID(
		cxt context.Context,
		id int,
	) (*sellermodel.Seller, error)
}

type getSellerBusiness struct {
	repo GetSellerRepo
}

func NewGetSellerBusiness(repo GetSellerRepo) *getSellerBusiness {
	return &getSellerBusiness{repo: repo}
}

func (biz *getSellerBusiness) GetSeller(
	ctx context.Context,
	id int,
	moreKeys ...string,
) (*sellermodel.Seller, error) {
	result, err := biz.repo.GetSellerWithID(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	return result, nil
}
