package cartmodel

import (
	"OpenMarket/common"

	"github.com/shopspring/decimal"
)

type CartItemDetailProduct struct {
	Id     int         `json:"-" gorm:"column:product_id"`
	FakeId *common.UID `json:"id" gorm:"-"`
	Name   string      `json:"name" gorm:"column:product_name"`
	Image  string      `json:"image" gorm:"column:image_url"`
}

func (p *CartItemDetailProduct) Mask() {
	p.FakeId = common.NewUID(uint32(p.Id), common.DbTypeProduct, 1)
}

type CartItemDetail struct {
	VariantId      int                    `json:"variant_id"`
	Quantity       int                    `json:"quantity"`
	Price          decimal.Decimal        `json:"price"`
	StockQuantity  int                    `json:"stock_quantity"`
	Product        CartItemDetailProduct  `json:"product"`
	Attributes     map[string]string      `json:"attributes"`
}

type CartItemDetailRow struct {
	VariantId      int              `gorm:"column:variant_id"`
	Quantity       int              `gorm:"column:quantity"`
	Price          decimal.Decimal  `gorm:"column:price"`
	StockQuantity  int              `gorm:"column:stock_quantity"`
	ProductId      int              `gorm:"column:product_id"`
	ProductName    string           `gorm:"column:product_name"`
	ImageURL       *string          `gorm:"column:image_url"`
	AttributeName  *string          `gorm:"column:attribute_name"`
	AttributeValue *string          `gorm:"column:attribute_value"`
}
