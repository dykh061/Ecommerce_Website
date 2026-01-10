package userstorage

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

func (s *sqlStore) FindAddressById(
	ctx context.Context,
	id, userId int,
) (*usermodel.UserAddress, error) {
	var data usermodel.UserAddress
	if err := s.db.Table(usermodel.UserAddress{}.TableName()).
		Where("id=? AND user_id=?", id, userId).
		First(&data).
		Error; err != nil {
		return nil, err
	}
	return &data, nil
}
