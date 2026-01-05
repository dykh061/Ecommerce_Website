package sellerrepository

import (
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type CreateSellerStorage interface {
	Create(ctx context.Context, data *sellermodel.SellerCreate) error
}

type createSellerRepo struct {
	storage CreateSellerStorage
}

func NewCreateSellerRepo(storage CreateSellerStorage) *createSellerRepo {
	return &createSellerRepo{
		storage: storage,
	}
}

func (repo *createSellerRepo) RegisterSeller(
	ctx context.Context,
	data *sellermodel.SellerCreate,
) error {
	return repo.storage.Create(ctx, data)
}
