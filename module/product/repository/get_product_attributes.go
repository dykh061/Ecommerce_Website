package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type GetProductAttributesStorage interface {
	GetProductAttributes(
		ctx context.Context,
		productID int,
	) ([]productmodel.AttributeValueRow, error)
}

type getProductAttributesRepo struct {
	storage GetProductAttributesStorage
}

func NewGetProductAttributesRepo(storage GetProductAttributesStorage) *getProductAttributesRepo {
	return &getProductAttributesRepo{storage: storage}
}

func (repo *getProductAttributesRepo) GetProductAttributes(
	ctx context.Context,
	productID int,
) ([]productmodel.ProductAttribute, error) {
	rows, err := repo.storage.GetProductAttributes(ctx, productID)
	if err != nil {
		return nil, err
	}

	// Group by attribute
	attrMap := make(map[int]*productmodel.ProductAttribute)
	order := make([]int, 0)

	for _, row := range rows {
		attr, exists := attrMap[row.AttributeID]
		if !exists {
			attr = &productmodel.ProductAttribute{
				ID:     row.AttributeID,
				Name:   row.AttributeName,
				Values: make([]productmodel.ProductAttributeValue, 0),
			}
			attrMap[row.AttributeID] = attr
			order = append(order, row.AttributeID)
		}
		attr.Values = append(attr.Values, productmodel.ProductAttributeValue{
			ID:    row.ValueID,
			Value: row.AttributeValue,
		})
	}

	// Build result in order
	result := make([]productmodel.ProductAttribute, 0, len(order))
	for _, id := range order {
		result = append(result, *attrMap[id])
	}

	return result, nil
}
