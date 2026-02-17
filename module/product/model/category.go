package productmodel

import "OpenMarket/common"

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"type:varchar(255);not null"`
	ParentId        int    `json:"parent_id" gorm:"not null;index"`
}

func (Category) TableName() string { return "categories" }

type CategoryCreate struct {
	Name     string `json:"name" form:"name" binding:"required"`
	ParentId int    `json:"parent_id" form:"parent_id"`
}

func (CategoryCreate) TableName() string { return Category{}.TableName() }

type CategoryUpdate struct {
	Name     *string `json:"name" gorm:"type:varchar(255)"`
	ParentId *int    `json:"parent_id"`
}

func (CategoryUpdate) TableName() string { return Category{}.TableName() }

type categoryListItem struct {
	Id       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	ParentId int    `json:"parent_id" gorm:"column:parent_id"`
}

func (categoryListItem) TableName() string { return Category{}.TableName() }
