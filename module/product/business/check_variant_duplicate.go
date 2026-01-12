package productbusiness

import (
	"OpenMarket/common"
	"context"
	"errors"
)

type VariantDuplicateRepo interface {
	CheckDuplicate(ctx context.Context, productID int, attributeValueIDs []int, excludeVariantID *int) (bool, error)
}

type checkVariantDuplicateBusiness struct {
	sfinder       FindSellerByID
	pfinder       FindProductWithIDAndSellerID
	duplicateRepo VariantDuplicateRepo
}

func NewCheckVariantDuplicateBusiness(
	sfinder FindSellerByID,
	pfinder FindProductWithIDAndSellerID,
	duplicateRepo VariantDuplicateRepo,
) *checkVariantDuplicateBusiness {
	return &checkVariantDuplicateBusiness{
		sfinder:       sfinder,
		pfinder:       pfinder,
		duplicateRepo: duplicateRepo,
	}
}

func (biz *checkVariantDuplicateBusiness) CheckDuplicate(
	ctx context.Context,
	userID int,
	productID int,
	attributeValueIDs []int,
	excludeVariantID *int,
) (bool, error) {
	// 1. Find seller
	seller, err := biz.sfinder.FindActiveSellerWithUserID(ctx, userID)
	if err != nil || seller == nil {
		return false, common.ErrForbidden(errors.New("user is not a seller"))
	}

	// 2. Check product ownership
	product, err := biz.pfinder.FindProductByIdWithSellerID(ctx, productID, seller.Id)
	if err != nil || product == nil {
		return false, common.ErrForbidden(errors.New("product does not belong to this seller"))
	}

	// 3. Check duplicate
	return biz.duplicateRepo.CheckDuplicate(ctx, productID, attributeValueIDs, excludeVariantID)
}
