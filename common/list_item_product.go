package common

type ListProductRequest struct {
	Page       int      `form:"page"`
	Limit      int      `form:"limit"`
	Search     *string  `form:"search"`
	CategoryID *int     `form:"category_id"`
	SellerID   *string  `form:"seller_id"`
	MinPrice   *float64 `form:"min_price"`
	MaxPrice   *float64 `form:"max_price"`
}
