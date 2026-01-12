package productmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

// VariantDetailFull is the response for GET /v1/seller/products/:id/variant/:vid
type VariantDetailFull struct {
	ID            int                       `json:"id"`
	Sku           string                    `json:"sku"`
	Price         decimal.Decimal           `json:"price"`
	StockQuantity int                       `json:"stock_quantity"`
	Status        int                       `json:"status"`
	Attributes    []VariantAttributeDetail  `json:"attributes"`
	CreatedAt     *time.Time                `json:"created_at,omitempty"`
	UpdatedAt     *time.Time                `json:"updated_at,omitempty"`
}

// VariantAttributeDetail includes attribute_id and attribute_value_id
type VariantAttributeDetail struct {
	AttributeID      int    `json:"attribute_id"`
	AttributeName    string `json:"attribute_name"`
	AttributeValueID int    `json:"attribute_value_id"`
	AttributeValue   string `json:"attribute_value"`
}

// VariantAttrFullRow is used for scanning variant with full attribute info
type VariantAttrFullRow struct {
	VariantID        int             `gorm:"column:variant_id"`
	Sku              string          `gorm:"column:sku"`
	Price            decimal.Decimal `gorm:"column:price"`
	StockQuantity    int             `gorm:"column:stock_quantity"`
	Status           int             `gorm:"column:status"`
	CreatedAt        *time.Time      `gorm:"column:created_at"`
	UpdatedAt        *time.Time      `gorm:"column:updated_at"`
	AttributeID      int             `gorm:"column:attribute_id"`
	AttributeName    string          `gorm:"column:attribute_name"`
	AttributeValueID int             `gorm:"column:attribute_value_id"`
	AttributeValue   string          `gorm:"column:attribute_value"`
}
