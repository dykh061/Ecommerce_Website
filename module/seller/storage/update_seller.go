package sellerstorage

import (
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

func (s *sqlStore) UpdateSeller(ctx context.Context, condition map[string]interface{}, data sellermodel.SellerUpdate) error {
	if err := s.db.Table(sellermodel.SellerUpdate{}.TableName()).
		Where(condition).
		Updates(data).Error; err != nil {
		return err
	}
	return nil
}
