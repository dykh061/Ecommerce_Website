package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
	"errors"
)

type GetCategoryAttributesRepo interface {
	CategoryExists(ctx context.Context, categoryID int) (bool, error)
	GetCategoryAttributes(ctx context.Context, categoryID int) ([]productmodel.CategoryAttributeWithValues, error)
}

type getCategoryAttributesBusiness struct {
	repo GetCategoryAttributesRepo
}

func NewGetCategoryAttributesBusiness(repo GetCategoryAttributesRepo) *getCategoryAttributesBusiness {
	return &getCategoryAttributesBusiness{repo: repo}
}

func (biz *getCategoryAttributesBusiness) GetCategoryAttributes(
	ctx context.Context,
	categoryID int,
) ([]productmodel.CategoryAttributeWithValues, error) {
	// Check category exists
	exists, err := biz.repo.CategoryExists(ctx, categoryID)
	if err != nil {
		return nil, common.ErrEntityNotFound("Category", err)
	}
	if !exists {
		return nil, common.ErrEntityNotFound("Category", errors.New("category not found"))
	}

	return biz.repo.GetCategoryAttributes(ctx, categoryID)
}
