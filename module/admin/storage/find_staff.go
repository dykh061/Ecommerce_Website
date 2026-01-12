package adminstorage

import (
	adminmodel "OpenMarket/module/admin/model"
	"context"
)

// FindStaffByUsername finds a staff by username with roles preloaded
func (s *sqlStore) FindStaffByUsername(
	ctx context.Context,
	username string,
) (*adminmodel.Staff, error) {
	var staff adminmodel.Staff
	if err := s.db.WithContext(ctx).
		Preload("Roles").
		Where("username = ?", username).
		First(&staff).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}

// FindStaffById finds a staff by id with roles preloaded
func (s *sqlStore) FindStaffById(
	ctx context.Context,
	id int,
) (*adminmodel.Staff, error) {
	var staff adminmodel.Staff
	if err := s.db.WithContext(ctx).
		Preload("Roles").
		Where("id = ?", id).
		First(&staff).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}

// CreateStaff creates a new staff account
func (s *sqlStore) CreateStaff(
	ctx context.Context,
	data *adminmodel.StaffCreate,
	hashedPassword string,
) (int, error) {
	staff := adminmodel.Staff{
		Username: data.Username,
		Password: hashedPassword,
		Status:   1,
	}
	if err := s.db.WithContext(ctx).Create(&staff).Error; err != nil {
		return 0, err
	}
	return staff.Id, nil
}

// AssignRolesToStaff assigns roles to a staff
func (s *sqlStore) AssignRolesToStaff(
	ctx context.Context,
	staffId int,
	roleIds []int,
) error {
	if len(roleIds) == 0 {
		return nil
	}
	staffRoles := make([]adminmodel.StaffRole, len(roleIds))
	for i, roleId := range roleIds {
		staffRoles[i] = adminmodel.StaffRole{
			StaffId: staffId,
			RoleId:  roleId,
		}
	}
	return s.db.WithContext(ctx).Create(&staffRoles).Error
}

// ListRoles returns all available roles
func (s *sqlStore) ListRoles(ctx context.Context) ([]adminmodel.Role, error) {
	var roles []adminmodel.Role
	if err := s.db.WithContext(ctx).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
