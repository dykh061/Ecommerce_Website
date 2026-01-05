package productmodel

import (
	"OpenMarket/common"
	"errors"

	"github.com/shopspring/decimal"
)

type Variant struct {
	common.SQLModel `json:",inline"`
	ProductId       int             `json:"-" gorm:"not null,index"`
	Sku             string          `json:"-" gorm:"type:varchar(100);not null"` //mã định danh nội bộ cho từng biến thể hàng hóa
	Price           decimal.Decimal `json:"price" gorm:"type:decimal(12,2);not null"`
	StockQuantity   int             `json:"stock_quantity" gorm:"not null"`
}

func (Variant) TableName() string { return "variants" }

type VariantCreate struct {
	common.SQLModel   `json:",inline"`
	ProductId         int             `json:"-" gorm:"not null,index"`
	Sku               string          `json:"-" gorm:"type:varchar(100);not null"` //mã định danh nội bộ cho từng biến thể hàng hóa
	Price             decimal.Decimal `json:"price" gorm:"type:decimal(12,2);not null"`
	StockQuantity     int             `json:"stock_quantity" gorm:"not null"`
	AttributeValueIDs []int           `json:"attribute_value_ids" gorm:"-"`
}

func (VariantCreate) TableName() string {
	return Variant{}.TableName()
}

type VariantUpdate struct {
	Price         *decimal.Decimal `json:"price" gorm:"type:decimal(12,2)"`
	StockQuantity *int             `json:"stock_quantity"`
}

func (VariantUpdate) TableName() string {
	return Variant{}.TableName()
}

func (v *VariantCreate) Validate() error {
	if v.Price.LessThanOrEqual(decimal.Zero) {
		return errors.New("price must be greater than 0")
	}
	return nil
}
