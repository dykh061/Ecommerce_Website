package sellerstorage

import (
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *sellermodel.SellerCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
