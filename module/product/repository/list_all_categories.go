package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type ListAllCategoriesStorage interface {
	ListAllCategories(ctx context.Context) ([]productmodel.CategoryListItem, error)
}

type listAllCategoriesRepo struct {
	storage ListAllCategoriesStorage
}

func NewListAllCategoriesRepo(storage ListAllCategoriesStorage) *listAllCategoriesRepo {
	return &listAllCategoriesRepo{storage: storage}
}

func (repo *listAllCategoriesRepo) ListAllCategories(
	ctx context.Context,
) ([]productmodel.CategoryListItem, error) {
	return repo.storage.ListAllCategories(ctx)
}
