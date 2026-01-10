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

type UserAddress struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"-" gorm:"column:user_id;index;not null"`
	Address         string `json:"address" gorm:"type:varchar(255);not null"`
	City            string `json:"city" gorm:"type:varchar(100);not null"`
}

func (UserAddress) TableName() string { return "user_addresses" }

type UserAddressCreate struct {
	UserId  int    `json:"-" gorm:"column:user_id;index;not null"`
	Address string `json:"address" gorm:"type:varchar(255);not null"`
	City    string `json:"city" gorm:"type:varchar(100);not null"`
}

func (UserAddressCreate) TableName() string { return UserAddress{}.TableName() }

type UserAddressUpdate struct {
	Address *string `json:"address" gorm:"type:varchar(255);not null"`
	Status  int     `json:"status" gorm:"column:status;default:0;index"`
	City    *string `json:"city" gorm:"type:varchar(100);not null"`
}

func (u *UserAddress) Mask() {
	u.GenUID(common.DbTypeAddress)
}

func (UserAddressUpdate) TableName() string { return UserAddress{}.TableName() }

func (UserLogin) TableName() string { return User{}.TableName() }

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetUserEmail() string {
	return u.Email
}
