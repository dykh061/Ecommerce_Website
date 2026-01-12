package productbusiness

import (
	"OpenMarket/common"
	"context"
	"errors"
)

type GalleryManagementRepo interface {
	SetMainGallery(ctx context.Context, productID int, galleryID int) error
	DeleteGallery(ctx context.Context, productID int, galleryID int) error
	GalleryExists(ctx context.Context, productID int, galleryID int) (bool, error)
}

type galleryManagementBusiness struct {
	sfinder     FindSellerByID
	pfinder     FindProductWithIDAndSellerID
	galleryRepo GalleryManagementRepo
}

func NewGalleryManagementBusiness(
	sfinder FindSellerByID,
	pfinder FindProductWithIDAndSellerID,
	galleryRepo GalleryManagementRepo,
) *galleryManagementBusiness {
	return &galleryManagementBusiness{
		sfinder:     sfinder,
		pfinder:     pfinder,
		galleryRepo: galleryRepo,
	}
}

func (biz *galleryManagementBusiness) SetMainGallery(
	ctx context.Context,
	userID int,
	productID int,
	galleryID int,
) error {
	// 1. Check seller ownership
	seller, err := biz.sfinder.FindActiveSellerWithUserID(ctx, userID)
	if err != nil || seller == nil {
		return common.ErrForbidden(errors.New("user is not a seller"))
	}

	// 2. Check product ownership
	product, err := biz.pfinder.FindProductByIdWithSellerID(ctx, productID, seller.Id)
	if err != nil || product == nil {
		return common.ErrForbidden(errors.New("product does not belong to this seller"))
	}

	// 3. Check gallery exists
	exists, err := biz.galleryRepo.GalleryExists(ctx, productID, galleryID)
	if err != nil {
		return common.ErrEntityNotFound("Gallery", err)
	}
	if !exists {
		return common.ErrEntityNotFound("Gallery", errors.New("gallery not found"))
	}

	// 4. Set main gallery
	if err := biz.galleryRepo.SetMainGallery(ctx, productID, galleryID); err != nil {
		return common.ErrCannotUpdateEntity("Gallery", err)
	}

	return nil
}

func (biz *galleryManagementBusiness) DeleteGallery(
	ctx context.Context,
	userID int,
	productID int,
	galleryID int,
) error {
	// 1. Check seller ownership
	seller, err := biz.sfinder.FindActiveSellerWithUserID(ctx, userID)
	if err != nil || seller == nil {
		return common.ErrForbidden(errors.New("user is not a seller"))
	}

	// 2. Check product ownership
	product, err := biz.pfinder.FindProductByIdWithSellerID(ctx, productID, seller.Id)
	if err != nil || product == nil {
		return common.ErrForbidden(errors.New("product does not belong to this seller"))
	}

	// 3. Check gallery exists
	exists, err := biz.galleryRepo.GalleryExists(ctx, productID, galleryID)
	if err != nil {
		return common.ErrEntityNotFound("Gallery", err)
	}
	if !exists {
		return common.ErrEntityNotFound("Gallery", errors.New("gallery not found"))
	}

	// 4. Delete gallery
	if err := biz.galleryRepo.DeleteGallery(ctx, productID, galleryID); err != nil {
		return common.ErrCannotDeleteEntity("Gallery", err)
	}

	return nil
}
