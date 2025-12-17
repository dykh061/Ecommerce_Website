package productmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID          int             `json:"id" gorm:"primaryKey;autoIncrement"`
	SellerID    int             `json:"seller_id" gorm:"not null;index"`
	Name        string          `json:"name" gorm:"type:varchar(255);not null"`
	Description string          `json:"description" gorm:"type:text"`
	BasePrice   decimal.Decimal `json:"base_price" gorm:"not null"`
	Status      string          `json:"status" gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime"`
}

func (Product) TableName() string { return "products" }

type ProductCreate struct {
	ID          int             `json:"id" gorm:"primaryKey;autoIncrement"`
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
	Status      *string          `json:"status" gorm:"type:varchar(20);not null"`
}

func (ProductUpdate) TableName() string { return Product{}.TableName() }
