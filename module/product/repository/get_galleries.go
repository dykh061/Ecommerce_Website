package productrepository

import (
	productmodel "OpenMarket/module/product/model"
	"context"
)

type GetGalleriesStorage interface {
	GetGalleries(
		ctx context.Context,
		productID int,
	) ([]productmodel.GalleryItem, error)
}

type getGalleriesRepo struct {
	storage GetGalleriesStorage
}

func NewGetGalleriesRepo(storage GetGalleriesStorage) *getGalleriesRepo {
	return &getGalleriesRepo{storage: storage}
}

func (repo *getGalleriesRepo) GetGalleries(
	ctx context.Context,
	productID int,
) ([]productmodel.GalleryItem, error) {
	return repo.storage.GetGalleries(ctx, productID)
}
