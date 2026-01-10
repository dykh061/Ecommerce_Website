package userstorage

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

func (s *sqlStore) UpdateAddress(
	ctx context.Context,
	userId, id int,
	data *usermodel.UserAddressUpdate,
) error {
	if err := s.db.Table(usermodel.UserAddressUpdate{}.TableName()).
		Where("user_id = ? AND id = ?", userId, id).
		Updates(*data).
		Error; err != nil {
		return err
	}
	return nil
}
