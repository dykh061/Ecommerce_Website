package usermodel

import (
	"OpenMarket/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"type:varchar(255);not null"`
	Email           string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password        string `json:"-" gorm:"type:varchar(255);not null"`
	Phone           string `json:"phone" gorm:"type:varchar(50)"`
	IsBanned        bool   `json:"is_banned" gorm:"type:boolean;default:false"`
}

func (User) TableName() string { return "users" }

func (u *User) Mask() {
	u.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Phone    string `json:"phone" gorm:"type:varchar(50)"`
}

func (UserCreate) TableName() string { return User{}.TableName() }

type UserUpdate struct {
	Name     *string `json:"name" gorm:"type:varchar(255);not null"`
	Email    *string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password *string `json:"password" gorm:"type:varchar(255);not null"`
	Phone    *string `json:"phone" gorm:"type:varchar(50)"`
	IsBanned *bool   `json:"is_banned" gorm:"type:boolean;default:false"`
}

func (UserUpdate) TableName() string { return User{}.TableName() }

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string { return User{}.TableName() }

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetUserEmail() string {
	return u.Email
}
