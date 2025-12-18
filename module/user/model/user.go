package usermodel

import (
	"time"
)

var (
	UserStatusActive  = "active"
	UserStatusDeleted = "deleted"
	UserStatusBanned  = "banned"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Phone     string    `json:"phone" gorm:"type:varchar(50)"`
	Status    string    `json:"status" gorm:"type:varchar(20);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (User) TableName() string { return "users" }

type UserCreate struct {
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Phone    string `json:"phone" gorm:"type:varchar(50)"`
	Status   string `json:"status" gorm:"type:varchar(20);not null"`
}

func (UserCreate) TableName() string { return User{}.TableName() }

type UserUpdate struct {
	Name     *string `json:"name" gorm:"type:varchar(255);not null"`
	Email    *string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password *string `json:"password" gorm:"type:varchar(255);not null"`
	Phone    *string `json:"phone" gorm:"type:varchar(50)"`
	Status   *string `json:"status" gorm:"type:varchar(20);not null"`
}

func (UserUpdate) TableName() string { return User{}.TableName() }
