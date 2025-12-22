package usermodel

import (
	"OpenMarket/common"
	"errors"
	"strings"
)

type User struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"type:varchar(255);not null"`
	Email           string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password        string `json:"-" gorm:"type:varchar(255);not null"`
	Phone           string `json:"phone" gorm:"type:varchar(50)"`
	IsBanned        bool   `json:"is_banned" gorm:"type:boolean;default:false"`
}

func (User) TableName() string { return "users" }

type UserCreate struct {
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Phone    string `json:"phone" gorm:"type:varchar(50)"`
}

func (data *UserCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
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

var (
	ErrNameIsEmpty = errors.New("Name can not be empty")
)
