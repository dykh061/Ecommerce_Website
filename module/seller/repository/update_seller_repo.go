package sellerrepository

import (
	"OpenMarket/common"
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type UpdateSellerStorage interface {
	UpdateSeller(ctx context.Context, condition map[string]interface{}, data sellermodel.SellerUpdate) error
}

type updateSellerRepo struct {
	storage UpdateSellerStorage
}

func NewUpdateSellerRepo(storage UpdateSellerStorage) *updateSellerRepo {
	return &updateSellerRepo{storage: storage}
}

func (repo *updateSellerRepo) UpdateSeller(
	ctx context.Context,
	userId int,
	data sellermodel.SellerUpdate,
) error {
	return repo.storage.UpdateSeller(ctx, map[string]interface{}{
		"user_id": userId,
		"status":  common.SystemStatusActive,
	}, data)
}
