package adminmodel

import (
	"OpenMarket/common"
	"strings"
	"time"
)

// Staff represents the staff_accounts table
type Staff struct {
	Id        int        `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	FakeId    *common.UID `json:"id" gorm:"-"`
	Username  string     `json:"username" gorm:"column:username;type:varchar(100);uniqueIndex;not null"`
	Password  string     `json:"-" gorm:"column:password;type:varchar(255);not null"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	Roles     []Role     `json:"roles,omitempty" gorm:"many2many:staff_roles;foreignKey:Id;joinForeignKey:staff_id;References:Id;joinReferences:role_id"`
}

func (Staff) TableName() string { return "staff_accounts" }

func (s *Staff) Mask() {
	s.FakeId = common.NewUID(uint32(s.Id), DbTypeStaff, 1)
}

func (s *Staff) GetStaffId() int {
	return s.Id
}

func (s *Staff) GetUsername() string {
	return s.Username
}

func (s *Staff) GetRoles() []string {
	roles := make([]string, len(s.Roles))
	for i, r := range s.Roles {
		roles[i] = r.Name
	}
	return roles
}

// HasRole checks if staff has a specific role
func (s *Staff) HasRole(roleName string) bool {
	for _, r := range s.Roles {
		if strings.EqualFold(strings.TrimSpace(r.Name), strings.TrimSpace(roleName)) {
			return true
		}
	}
	return false
}

// StaffCreate for creating new staff
type StaffCreate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	RoleIds  []int  `json:"role_ids"`
}

func (StaffCreate) TableName() string { return Staff{}.TableName() }

// StaffLogin for staff authentication
type StaffLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// StaffUpdate for updating staff info
type StaffUpdate struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
	Status   *int    `json:"status"`
}

func (StaffUpdate) TableName() string { return Staff{}.TableName() }

const (
	DbTypeStaff = 10
)
