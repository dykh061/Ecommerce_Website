package sellerstorage

import (
	"OpenMarket/common"
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

func (s *sqlStore) ListSellers(
	ctx context.Context,
	filter *sellermodel.SellerFilter,
	paging *common.Paging,
	moreKeys ...string,
) ([]sellermodel.Seller, error) {
	var sellers []sellermodel.Seller
	db := s.db.Table(sellermodel.Seller{}.TableName()).
		Where("status = ?", common.SystemStatusActive)

	if f := filter; f != nil {
		if f.Id != nil {
			db = db.Where("id = ?", *f.Id)
		}
		if f.Keyword != nil {
			db = db.Where(
				"LOWER(shop_name) LIKE LOWER(?) OR LOWER(shop_description) LIKE LOWER(?)",
				"%"+*f.Keyword+"%",
				"%"+*f.Keyword+"%",
			)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	if err := db.Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&sellers).
		Error; err != nil {
		return nil, err
	}
	return sellers, nil
}
