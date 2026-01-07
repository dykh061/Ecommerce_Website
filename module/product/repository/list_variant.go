package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type ListVariantStorage interface {
	ListVariant(
		ctx context.Context,
		productID int,
	) ([]productmodel.VariantAttrRow, error)
}

type listVariantRepo struct {
	storage ListVariantStorage
}

func NewListVariantRepo(storage ListVariantStorage) *listVariantRepo {
	return &listVariantRepo{storage: storage}
}
func (repo *listVariantRepo) ListVariant(
	ctx context.Context,
	productID int,
) ([]productmodel.VariantDetail, error) {
	rows, err := repo.storage.ListVariant(ctx, productID)
	if err != nil {
		return nil, err
	}
	variantMap := make(map[int]*productmodel.VariantDetail)
	for _, row := range rows {
		v, ok := variantMap[row.VariantID]
		if !ok {
			v = &productmodel.VariantDetail{
				ID:            row.VariantID,
				Sku:           row.Sku,
				Price:         row.Price,
				StockQuantity: row.StockQuantity,
			}
			variantMap[row.VariantID] = v
		}
		v.Attributes = append(v.Attributes, productmodel.VariantAttribute{
			AttributeID:   row.AttributeID,
			AttributeName: row.AttributeName,
			Value:         row.AttributeValue,
		})
	}
	result := make([]productmodel.VariantDetail, 0, len(variantMap))
	for _, v := range variantMap {
		result = append(result, *v)
	}

	return result, nil
}
