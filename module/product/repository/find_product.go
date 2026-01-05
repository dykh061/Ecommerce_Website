package productrepository

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

type FindProductStorage interface {
	FindProductWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
}

type findProductRepo struct {
	storage FindProductStorage
}

func NewFindProductRepo(storage FindProductStorage) *findProductRepo {
	return &findProductRepo{storage: storage}
}

func (repo *findProductRepo) FindProductByIdWithSellerID(
	ctx context.Context,
	id int,
	sellerID int,
) (*productmodel.Product, error) {
	return repo.storage.FindProductWithCondition(ctx, map[string]interface{}{
		"id":        id,
		"seller_id": sellerID,
		"status":    common.SystemStatusActive,
	})
}
