package ordermodel

import "github.com/shopspring/decimal"

type OrderItem struct {
	Id        int             `json:"id" gorm:"column:id"`
	OrderId   int             `json:"order_id" gorm:"column:order_id"`
	VariantId int             `json:"variant_id" gorm:"column:variant_id"`
	Quantity  int             `json:"quantity" gorm:"column:quantity"`
	Price     decimal.Decimal `json:"price" gorm:"column:price"`
}

func (OrderItem) TableName() string {
	return "order_items"
}

type OrderItemCreate struct {
	OrderId   int             `json:"-" gorm:"column:order_id;not null;index"`
	VariantId int             `json:"variant_id" gorm:"column:variant_id;not null"`
	Quantity  int             `json:"quantity" gorm:"column:quantity;not null"`
	Price     decimal.Decimal `json:"price" gorm:"column:price;type:decimal(12,2);not null"`
}

func (OrderItemCreate) TableName() string {
	return OrderItem{}.TableName()
}

type OrderItemView struct {
	VariantId int             `json:"variant_id"`
	Quantity  int             `json:"quantity"`
	Price     decimal.Decimal `json:"price"`
	SubTotal  decimal.Decimal `json:"sub_total"`
}
