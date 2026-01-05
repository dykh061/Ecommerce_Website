package sellerrepository

import (
	"OpenMarket/common"
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type GetSellerStorage interface {
	FindSellerWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*sellermodel.Seller, error)
}

type getSellerRepo struct {
	storage GetSellerStorage
}

func NewGetSellerRepo(storage GetSellerStorage) *getSellerRepo {
	return &getSellerRepo{storage: storage}
}

func (repo *getSellerRepo) FindActiveSellerWithUserID(
	ctx context.Context,
	userid int,
) (*sellermodel.Seller, error) {
	return repo.storage.FindSellerWithCondition(ctx, map[string]interface{}{
		"user_id": userid,
		"status":  common.SystemStatusActive,
	})
}
