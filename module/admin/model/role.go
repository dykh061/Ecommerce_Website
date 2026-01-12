package adminmodel

// Role represents the roles table
type Role struct {
	Id   int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"column:name;type:varchar(50);uniqueIndex;not null"`
}

func (Role) TableName() string { return "roles" }

// Predefined role names
const (
	RoleAdmin      = "admin"
	RoleModerator  = "moderator"
	RoleSupport    = "support"
)

// StaffRole represents the staff_roles junction table
type StaffRole struct {
	StaffId int `gorm:"column:staff_id;primaryKey"`
	RoleId  int `gorm:"column:role_id;primaryKey"`
}

func (StaffRole) TableName() string { return "staff_roles" }
