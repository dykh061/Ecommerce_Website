package productbusiness

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type DeleteProductStore interface {
	FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*productmodel.Product, error)
	Delete(ctx context.Context, id int) error
}

type deleteProductBusiness struct {
	store DeleteProductStore
}

func NewDeleteProductBusiness(store DeleteProductStore) *deleteProductBusiness {
	return &deleteProductBusiness{store: store}
}

// func (biz *deleteProductBusiness) DeleteProduct(ctx context.Context, id int) error {
// 	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
// 	if err != nil {
// 		return err
// 	}
// 	if oldData.Status == "deleted" {
// 		return errors.New("Data has been deleted")
// 	}
// 	if err := biz.store.Delete(ctx, id); err != nil {
// 		return err
// 	}
// 	return nil
// }
