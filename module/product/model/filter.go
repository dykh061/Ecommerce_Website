package productmodel

import "github.com/shopspring/decimal"

type Filter struct {
	SellerID   *int             `json:"seller_id" form:"seller_id"`
	Status     *string          `json:"status" form:"status"`
	CategoryID *int             `json:"category_id" form:"category_id"`
	MinPrice   *decimal.Decimal `json:"min_price" form:"min_price"`
	MaxPrice   *decimal.Decimal `json:"max_price" form:"max_price"`
	Search     *string          `json:"search" form:"search"`
}
