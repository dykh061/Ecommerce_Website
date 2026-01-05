package sellerstorage

import (
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

func (s *sqlStore) FindSellerWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	morekeys ...string,
) (*sellermodel.Seller, error) {
	db := s.db.Table(sellermodel.Seller{}.TableName())

	for i := range morekeys {
		db = db.Preload(morekeys[i])
	}
	var data sellermodel.Seller

	if err := db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
