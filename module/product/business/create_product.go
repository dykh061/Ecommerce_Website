package productbusiness

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type CreateProductStore interface {
	Create(ctx context.Context, data *productmodel.ProductCreate) error
}

type createProductBusiness struct {
	store CreateProductStore
}

func NewCreateProductBusiness(store CreateProductStore) *createProductBusiness {
	return &createProductBusiness{store: store}
}

func (biz *createProductBusiness) CreateProduct(ctx context.Context, data *productmodel.ProductCreate) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
