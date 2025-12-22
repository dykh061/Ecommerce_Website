package userstorage

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

func (s *sqlStore) Update(ctx context.Context, condition map[string]interface{}, data usermodel.UserUpdate) error {
	if err := s.db.Table(usermodel.UserUpdate{}.TableName()).
		Where(condition).
		Updates(data).
		Error; err != nil {
		return err
	}
	return nil
}
