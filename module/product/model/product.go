package productmodel

import (
	"OpenMarket/common"

	"github.com/shopspring/decimal"
)

type Product struct {
	common.SQLModel `json:",inline"`
	SellerID        int             `json:"seller_id" gorm:"not null;index"`
	Name            string          `json:"name" gorm:"type:varchar(255);not null"`
	Description     string          `json:"description" gorm:"type:text"`
	BasePrice       decimal.Decimal `json:"base_price" gorm:"not null"`
}

func (Product) TableName() string { return "products" }

type ProductCreate struct {
	SellerID    int             `json:"seller_id" gorm:"not null"`
	Name        string          `json:"name" gorm:"type:varchar(255);not null"`
	Description string          `json:"description" gorm:"type:text"`
	BasePrice   decimal.Decimal `json:"base_price" gorm:"not null"`
}

func (ProductCreate) TableName() string { return Product{}.TableName() }

type ProductUpdate struct {
	Name        *string          `json:"name" gorm:"type:varchar(255);not null"`
	Description *string          `json:"description" gorm:"type:text"`
	BasePrice   *decimal.Decimal `json:"base_price" gorm:"not null"`
}

func (ProductUpdate) TableName() string { return Product{}.TableName() }
