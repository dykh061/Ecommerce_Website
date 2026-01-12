package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type GetVariantWithAttrsStorage interface {
	GetVariantWithAttributes(
		ctx context.Context,
		variantID int,
		productID int,
	) ([]productmodel.VariantAttrFullRow, error)
}

type getVariantWithAttrsRepo struct {
	storage GetVariantWithAttrsStorage
}

func NewGetVariantWithAttrsRepo(storage GetVariantWithAttrsStorage) *getVariantWithAttrsRepo {
	return &getVariantWithAttrsRepo{storage: storage}
}

func (repo *getVariantWithAttrsRepo) GetVariantWithAttributes(
	ctx context.Context,
	variantID int,
	productID int,
) (*productmodel.VariantDetailFull, error) {
	rows, err := repo.storage.GetVariantWithAttributes(ctx, variantID, productID)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}

	// First row has the variant info
	firstRow := rows[0]
	result := &productmodel.VariantDetailFull{
		ID:            firstRow.VariantID,
		Sku:           firstRow.Sku,
		Price:         firstRow.Price,
		StockQuantity: firstRow.StockQuantity,
		Status:        firstRow.Status,
		CreatedAt:     firstRow.CreatedAt,
		UpdatedAt:     firstRow.UpdatedAt,
		Attributes:    make([]productmodel.VariantAttributeDetail, 0),
	}

	// Collect attributes
	for _, row := range rows {
		if row.AttributeID > 0 {
			result.Attributes = append(result.Attributes, productmodel.VariantAttributeDetail{
				AttributeID:      row.AttributeID,
				AttributeName:    row.AttributeName,
				AttributeValueID: row.AttributeValueID,
				AttributeValue:   row.AttributeValue,
			})
		}
	}

	return result, nil
}
