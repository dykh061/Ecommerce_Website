package productbusiness

import "context"

type getImagesBusiness struct {
	repo GetImagesRepo
}

func NewGetImagesBusiness(repo GetImagesRepo) *getImagesBusiness {
	return &getImagesBusiness{repo: repo}
}

func (business *getImagesBusiness) GetImages(
	ctx context.Context,
	productID int,
) ([]string, error) {
	return business.repo.GetImages(ctx, productID)
}
