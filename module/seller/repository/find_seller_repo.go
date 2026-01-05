package sellerrepository

import (
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type FindSellerStorage interface {
	FindSellerWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*sellermodel.Seller, error)
}

type findSellerRepo struct {
	storage FindSellerStorage
}

func NewFindSellerRepo(storage FindSellerStorage) *findSellerRepo {
	return &findSellerRepo{storage: storage}
}

func (repo *findSellerRepo) FindSeller(
	ctx context.Context,
	userid int,
) (*sellermodel.Seller, error) {
	return repo.storage.FindSellerWithCondition(ctx, map[string]interface{}{"user_id": userid})
}
