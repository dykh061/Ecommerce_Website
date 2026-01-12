package adminstorage

import (
	adminmodel "OpenMarket/module/admin/model"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) UpdateStaffPasswordAndEnable(
	ctx context.Context,
	staffId int,
	hashedPassword string,
) error {
	// Use UpdateColumns to avoid GORM auto-writing updated_at
	// (some DB schemas may not have updated_at column).
	return s.db.WithContext(ctx).
		Model(&adminmodel.Staff{}).
		Where("id = ?", staffId).
		UpdateColumns(map[string]any{"password": hashedPassword, "status": 1}).Error
}

func (s *sqlStore) EnsureStaffRole(
	ctx context.Context,
	staffId int,
	roleId int,
) error {
	var existing adminmodel.StaffRole
	err := s.db.WithContext(ctx).
		Where("staff_id = ? AND role_id = ?", staffId, roleId).
		First(&existing).Error
	if err == nil {
		return nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return s.db.WithContext(ctx).
		Create(&adminmodel.StaffRole{StaffId: staffId, RoleId: roleId}).Error
}
