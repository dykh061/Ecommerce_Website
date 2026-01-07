package productmodel

import (
	"github.com/shopspring/decimal"
)

type ProductDetail struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	BasePrice   decimal.Decimal `json:"base_price"`
	Images      []string        `json:"images"`
	Variants    []VariantDetail `json:"variants"`
}

type VariantDetail struct {
	ID            int                `json:"id"`
	Sku           string             `json:"sku"`
	Price         decimal.Decimal    `json:"price"`
	StockQuantity int                `json:"stock_quantity"`
	Attributes    []VariantAttribute `json:"attributes"`
}

type VariantAttribute struct {
	AttributeID   int    `json:"-"`
	AttributeName string `json:"name"`
	Value         string `json:"value"`
}

type VariantAttrRow struct {
	VariantID      int             `gorm:"column:variant_id" json:"-"`
	Sku            string          `gorm:"column:sku" json:"-"`
	Price          decimal.Decimal `gorm:"column:price" json:"-"`
	StockQuantity  int             `gorm:"column:stock_quantity" json:"-"`
	AttributeID    int             `gorm:"column:attribute_id" json:"-"`
	AttributeName  string          `gorm:"column:attribute_name" json:"-"`
	AttributeValue string          `gorm:"column:attribute_value" json:"-"`
}
