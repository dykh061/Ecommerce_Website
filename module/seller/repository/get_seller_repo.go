package sellerrepository

import (
	"OpenMarket/common"
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type GetSellerWithIdStorage interface {
	FindSellerWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*sellermodel.Seller, error)
}

type getSellerWithIdRepo struct {
	storage GetSellerWithIdStorage
}

func NewGetSellerWithIdRepo(storage GetSellerWithIdStorage) *getSellerWithIdRepo {
	return &getSellerWithIdRepo{
		storage: storage,
	}
}

func (repo *getSellerWithIdRepo) GetSellerWithID(
	cxt context.Context,
	id int,
) (*sellermodel.Seller, error) {
	return repo.storage.FindSellerWithCondition(cxt,
		map[string]interface{}{
			"id":     id,
			"status": common.SystemStatusActive,
		})
}
