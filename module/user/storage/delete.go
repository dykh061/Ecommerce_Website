package userstorage

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(usermodel.User{}.TableName()).
		Where("id=? AND status = ?", id, common.SystemStatusActive).
		Updates(map[string]interface{}{
			"status": common.SystemStatusDeleted,
		}).Error; err != nil {
		return err
	}
	return nil
}
