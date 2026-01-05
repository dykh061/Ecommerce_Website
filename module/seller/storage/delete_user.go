package sellerstorage

import (
	"OpenMarket/common"
	sellermodel "OpenMarket/module/seller/model"
	"context"
)

func (s *sqlStore) DeleteSeller(ctx context.Context, userId int) error {
	if err := s.db.Table(sellermodel.Seller{}.TableName()).
		Where("user_id = ? AND status = ?", userId, common.SystemStatusActive).
		Updates(map[string]interface{}{
			"status": common.SystemStatusDeleted,
		}).Error; err != nil {
		return err
	}
	return nil
}
