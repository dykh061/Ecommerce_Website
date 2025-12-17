package userstorage

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(usermodel.User{}.TableName()).
		Where("id=?", id).
		Updates(map[string]interface{}{"status": "deleted"}).
		Error; err != nil {
		return err
	}
	return nil
}
