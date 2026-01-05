package common

type SimpleUser struct {
	SQLModel `json:",inline"`
	Name     string `json:"name" gorm:"column:name;"`
}

func (SimpleUser) TableName() string {
	return "users"
}
func (u *SimpleUser) Mask() {
	u.GenUID(DbTypeUser)
}
