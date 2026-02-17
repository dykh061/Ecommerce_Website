package productmodel

import (
	"OpenMarket/common"

	"github.com/shopspring/decimal"
)

type Product struct {
	common.SQLModel `json:",inline"`
	SellerID        int             `json:"-" gorm:"not null;index"`
	Name            string          `json:"name" gorm:"type:varchar(255);not null"`
	Description     string          `json:"description" gorm:"type:text"`
	BasePrice       decimal.Decimal `json:"base_price" gorm:"not null"`
	CategoryID      *int            `json:"category_id" gorm:"index"`
}

func (Product) TableName() string { return "products" }

type ProductCreate struct {
	common.SQLModel `json:",inline"`
	SellerID        int             `json:"-" gorm:"not null"`
	Name            string          `json:"name" form:"name" binding:"required"`
	Description     string          `json:"description" form:"description"`
	BasePrice       decimal.Decimal `json:"base_price" form:"base_price" binding:"required"`
	CategoryID      *int            `json:"category_id" gorm:"index"`
}

func (ProductCreate) TableName() string { return Product{}.TableName() }

func (p *ProductCreate) Mask() {
	p.GenUID(common.DbTypeProduct)
}

type ProductUpdate struct {
	Name        *string          `json:"name" gorm:"type:varchar(255);not null"`
	Description *string          `json:"description" gorm:"type:text"`
	BasePrice   *decimal.Decimal `json:"base_price" gorm:"not null"`
	CategoryID  *int             `json:"category_id" gorm:"column:category_id"`
}

func (ProductUpdate) TableName() string { return Product{}.TableName() }

const (
	EntityName = "Product"
)

type ProductListItem struct {
	common.SQLModel
	Name      string          `json:"name"`
	BasePrice decimal.Decimal `json:"base_price"`
	ImageURL  *string         `json:"image_url,omitempty"`
}

func (p *ProductListItem) Mask() {
	p.GenUID(common.DbTypeProduct)
}
