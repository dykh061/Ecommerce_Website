package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type GetCategoryAttributesStorage interface {
	GetCategoryAttributes(ctx context.Context, categoryID int) ([]productmodel.CategoryAttributeRow, error)
	CategoryExists(ctx context.Context, categoryID int) (bool, error)
}

type getCategoryAttributesRepo struct {
	storage GetCategoryAttributesStorage
}

func NewGetCategoryAttributesRepo(storage GetCategoryAttributesStorage) *getCategoryAttributesRepo {
	return &getCategoryAttributesRepo{storage: storage}
}

func (repo *getCategoryAttributesRepo) CategoryExists(
	ctx context.Context,
	categoryID int,
) (bool, error) {
	return repo.storage.CategoryExists(ctx, categoryID)
}

func (repo *getCategoryAttributesRepo) GetCategoryAttributes(
	ctx context.Context,
	categoryID int,
) ([]productmodel.CategoryAttributeWithValues, error) {
	rows, err := repo.storage.GetCategoryAttributes(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	// Group by attribute
	attrMap := make(map[int]*productmodel.CategoryAttributeWithValues)
	order := make([]int, 0)

	for _, row := range rows {
		attr, exists := attrMap[row.AttributeID]
		if !exists {
			attr = &productmodel.CategoryAttributeWithValues{
				ID:     row.AttributeID,
				Name:   row.AttributeName,
				Values: make([]productmodel.AttributeValue, 0),
			}
			attrMap[row.AttributeID] = attr
			order = append(order, row.AttributeID)
		}
		if row.ValueID > 0 {
			attr.Values = append(attr.Values, productmodel.AttributeValue{
				ID:    row.ValueID,
				Value: row.AttributeValue,
			})
		}
	}

	// Build result in order
	result := make([]productmodel.CategoryAttributeWithValues, 0, len(order))
	for _, id := range order {
		result = append(result, *attrMap[id])
	}

	return result, nil
}
