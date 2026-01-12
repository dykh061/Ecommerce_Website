package productbusiness

import (
	"OpenMarket/common"
	productmodel "OpenMarket/module/product/model"
	"context"
	"errors"
)

type UpdateVariantStatusRepo interface {
	UpdateStatus(ctx context.Context, variantID int, productID int, status int) error
}

type VariantReader interface {
	FindVariantByID(ctx context.Context, id int) (*productmodel.Variant, error)
}

type toggleVariantStatusBusiness struct {
	sfinder       FindSellerByID
	pfinder       FindProductWithIDAndSellerID
	statusRepo    UpdateVariantStatusRepo
	variantReader VariantReader
}

func NewToggleVariantStatusBusiness(
	sfinder FindSellerByID,
	pfinder FindProductWithIDAndSellerID,
	statusRepo UpdateVariantStatusRepo,
	variantReader VariantReader,
) *toggleVariantStatusBusiness {
	return &toggleVariantStatusBusiness{
		sfinder:       sfinder,
		pfinder:       pfinder,
		statusRepo:    statusRepo,
		variantReader: variantReader,
	}
}

func (biz *toggleVariantStatusBusiness) ToggleStatus(
	ctx context.Context,
	userID int,
	productID int,
	variantID int,
	status int,
) error {
	// Validate status
	if status != 0 && status != 1 {
		return common.InvalidRequestError(errors.New("status must be 0 or 1"))
	}

	// 1. Find seller
	seller, err := biz.sfinder.FindActiveSellerWithUserID(ctx, userID)
	if err != nil || seller == nil {
		return common.ErrForbidden(errors.New("user is not a seller"))
	}

	// 2. Check product ownership
	product, err := biz.pfinder.FindProductByIdWithSellerID(ctx, productID, seller.Id)
	if err != nil || product == nil {
		return common.ErrForbidden(errors.New("product does not belong to this seller"))
	}

	// 3. Check variant exists
	variant, err := biz.variantReader.FindVariantByID(ctx, variantID)
	if err != nil || variant == nil {
		return common.ErrEntityNotFound("Variant", errors.New("variant not found"))
	}

	// 4. Update status
	if err := biz.statusRepo.UpdateStatus(ctx, variantID, productID, status); err != nil {
		return common.ErrCannotUpdateEntity("Variant", err)
	}

	return nil
}
