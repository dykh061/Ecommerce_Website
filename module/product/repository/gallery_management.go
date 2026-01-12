package productrepository

import "context"

type GalleryManagementStorage interface {
	SetMainGallery(ctx context.Context, productID int, galleryID int) error
	DeleteGallery(ctx context.Context, productID int, galleryID int) error
	GalleryExists(ctx context.Context, productID int, galleryID int) (bool, error)
}

type galleryManagementRepo struct {
	storage GalleryManagementStorage
}

func NewGalleryManagementRepo(storage GalleryManagementStorage) *galleryManagementRepo {
	return &galleryManagementRepo{storage: storage}
}

func (repo *galleryManagementRepo) SetMainGallery(
	ctx context.Context,
	productID int,
	galleryID int,
) error {
	return repo.storage.SetMainGallery(ctx, productID, galleryID)
}

func (repo *galleryManagementRepo) DeleteGallery(
	ctx context.Context,
	productID int,
	galleryID int,
) error {
	return repo.storage.DeleteGallery(ctx, productID, galleryID)
}

func (repo *galleryManagementRepo) GalleryExists(
	ctx context.Context,
	productID int,
	galleryID int,
) (bool, error) {
	return repo.storage.GalleryExists(ctx, productID, galleryID)
}
