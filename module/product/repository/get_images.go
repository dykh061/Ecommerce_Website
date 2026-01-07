package productrepository

import "context"

type GetImagesStorage interface {
	GetImages(
		context context.Context,
		productID int,
	) ([]string, error)
}
type getImagesRepo struct {
	storage GetImagesStorage
}

func NewGetImagesRepo(storage GetImagesStorage) *getImagesRepo {
	return &getImagesRepo{storage: storage}
}

func (repo *getImagesRepo) GetImages(
	ctx context.Context,
	productID int,
) ([]string, error) {
	return repo.storage.GetImages(ctx, productID)
}
