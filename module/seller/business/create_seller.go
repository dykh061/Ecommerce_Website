package sellerbusiness

import (
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type CreateSellerStore interface {
	Create(ctx context.Context, data *sellermodel.SellerCreate) error
}

type createSellerStore struct {
	store CreateSellerStore
}

func NewCreateSellerBusiness(store CreateSellerStore) *createSellerStore {
	return &createSellerStore{store: store}
}

func (biz *createSellerStore) CreateSeller(ctx context.Context, data *sellermodel.SellerCreate) error {

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
