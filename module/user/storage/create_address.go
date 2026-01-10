package userstorage

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

func (s *sqlStore) CreateAddress(
	ctx context.Context,
	create *usermodel.UserAddressCreate,
) error {
	if err := s.db.Create(&create).Error; err != nil {
		return err
	}
	return nil
}
