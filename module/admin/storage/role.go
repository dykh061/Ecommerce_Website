package adminstorage

import (
	adminmodel "OpenMarket/module/admin/model"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) FindRoleByName(
	ctx context.Context,
	name string,
) (*adminmodel.Role, error) {
	var role adminmodel.Role
	if err := s.db.WithContext(ctx).
		Where("name = ?", name).
		First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *sqlStore) CreateRole(
	ctx context.Context,
	name string,
) (int, error) {
	role := adminmodel.Role{Name: name}
	if err := s.db.WithContext(ctx).Create(&role).Error; err != nil {
		return 0, err
	}
	return role.Id, nil
}

func (s *sqlStore) EnsureRole(
	ctx context.Context,
	name string,
) (int, error) {
	role, err := s.FindRoleByName(ctx, name)
	if err == nil {
		return role.Id, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return s.CreateRole(ctx, name)
}
