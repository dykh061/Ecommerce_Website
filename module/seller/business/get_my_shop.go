package sellerbusiness

import (
	"OpenMarket/common"
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

type getMyShopBusiness struct {
	finder FindSellerByID
}

func NewGetMyShopBusiness(finder FindSellerByID) *getMyShopBusiness {
	return &getMyShopBusiness{finder: finder}
}

func (biz *getMyShopBusiness) GetMyShop(
	cxt context.Context,
	userID int,
	moreKeys ...string,
) (*sellermodel.Seller, error) {
	result, err := biz.finder.FindActiveSellerWithUserID(cxt, userID)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	if result != nil {
		if result.Status == common.SystemStatusDeleted {
			return nil, common.ErrInvalidState(sellermodel.EntityName, "no permission")
		}
	}
	return result, nil
}
