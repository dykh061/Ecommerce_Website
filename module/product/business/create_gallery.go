package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
)

type CreateGalleryRepo interface {
	CreateProductGallery(
		ctx context.Context,
		data *productmodel.GalleryCreate,
	) error
}

type createGalleryBiz struct {
	uploader ImageUploader
	repo     CreateGalleryRepo
	sfinder  FindSellerByID
	pfinder  FindProductWithIDAndSellerID
}

func NewCreateGalleryBiz(
	uploader ImageUploader,
	repo CreateGalleryRepo,
	sfinder FindSellerByID,
	pfinder FindProductWithIDAndSellerID,
) *createGalleryBiz {
	return &createGalleryBiz{
		uploader: uploader,
		repo:     repo,
		sfinder:  sfinder,
		pfinder:  pfinder,
	}
}

func (biz *createGalleryBiz) CreateProductGallery(
	ctx context.Context,
	productId int,
	userId int,
	fileBytes []byte,
	filename string,
	isMain bool,
) error {

	seller, err := biz.sfinder.FindActiveSellerWithUserID(ctx, userId)
	if err != nil {
		return common.InvalidRequestError(err)
	}
	_, err = biz.pfinder.FindProductByIdWithSellerID(ctx, productId, seller.Id)
	if err != nil {
		return common.InvalidRequestError(err)
	}

	img, err := biz.uploader.UploadFile(ctx, fileBytes, "product", filename)
	if err != nil {
		return common.InvalidRequestError(err)
	}
	imgDb := &productmodel.GalleryCreate{
		ProductId: productId,
		ImageURL:  img.Url,
		IsMain:    isMain,
	}
	if err := biz.repo.CreateProductGallery(ctx, imgDb); err != nil {
		return common.InvalidRequestError(err)
	}
	return nil
}
