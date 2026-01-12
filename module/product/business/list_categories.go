package productbusiness

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type ListAllCategoriesRepo interface {
	ListAllCategories(ctx context.Context) ([]productmodel.CategoryListItem, error)
}

type listCategoriesBusiness struct {
	repo ListAllCategoriesRepo
}

func NewListCategoriesBusiness(repo ListAllCategoriesRepo) *listCategoriesBusiness {
	return &listCategoriesBusiness{repo: repo}
}

func (biz *listCategoriesBusiness) ListCategories(
	ctx context.Context,
) ([]productmodel.CategoryListItem, error) {
	return biz.repo.ListAllCategories(ctx)
}
