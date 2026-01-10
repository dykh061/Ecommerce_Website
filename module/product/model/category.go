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

type ProductCategory struct {
	ProductId  int `json:"product_id" gorm:"not null;index"`
	CategoryId int `json:"category_id" gorm:"not null;index"`
}

func (ProductCategory) TableName() string { return "product_categories" }

type ProductCategoryUpdate struct {
	CategoryId *int `json:"category_id" gorm:"not null;index"`
}

func (ProductCategoryUpdate) TableName() string { return ProductCategory{}.TableName() }
