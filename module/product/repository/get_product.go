package productrepository

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

type GetProductStorage interface {
	FindProductWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
}

type getProductRepo struct {
	storage GetProductStorage
}

func NewGetProductRepo(storage GetProductStorage) *getProductRepo {
	return &getProductRepo{storage: storage}
}

func (repo *getProductRepo) GetProductById(
	ctx context.Context,
	productId int,
) (*productmodel.Product, error) {
	return repo.storage.FindProductWithCondition(ctx, map[string]interface{}{
		"id":     productId,
		"status": common.SystemStatusActive,
	})
}
