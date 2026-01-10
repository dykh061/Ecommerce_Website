package userstorage

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

func (s *sqlStore) GetListAddress(
	ctx context.Context,
	userId int,
) ([]usermodel.UserAddress, error) {
	var result []usermodel.UserAddress
	if err := s.db.Table(usermodel.UserAddress{}.TableName()).
		Where("user_id = ?", userId).
		Find(&result).
		Error; err != nil {
		return nil, err
	}
	return result, nil
}
