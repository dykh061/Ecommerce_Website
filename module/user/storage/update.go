package userstorage

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

func (s *sqlStore) Update(ctx context.Context, id int, data usermodel.UserUpdate) error {
	if err := s.db.Table(usermodel.UserUpdate{}.TableName()).
		Where("id=?", id).
		Updates(data).
		Error; err != nil {
		return err
	}
	return nil
}
