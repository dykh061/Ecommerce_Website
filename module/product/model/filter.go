package productmodel

import "github.com/shopspring/decimal"

type Filter struct {
	SellerID   *int             `json:"seller_id" form:"-" gorm:"seller_id"`
	Status     *int             `json:"status" form:"status" gorm:"status"`
	CategoryID *int             `json:"category_id" form:"category_id"`
	MinPrice   *decimal.Decimal `json:"min_price" form:"min_price"`
	MaxPrice   *decimal.Decimal `json:"max_price" form:"max_price"`
	Search     *string          `json:"search" form:"search"`
}
