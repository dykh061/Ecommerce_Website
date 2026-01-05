package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type CreateGalleryStore interface {
	CreateProductGallery(
		ctx context.Context,
		data *productmodel.GalleryCreate,
	) error
}

type createGalleryRepo struct {
	storage CreateGalleryStore
}

func NewCreateGalleryRepo(storage CreateGalleryStore) *createGalleryRepo {
	return &createGalleryRepo{storage: storage}
}

func (repo *createGalleryRepo) CreateProductGallery(
	ctx context.Context,
	data *productmodel.GalleryCreate,
) error {
	return repo.storage.CreateProductGallery(ctx, data)
}
