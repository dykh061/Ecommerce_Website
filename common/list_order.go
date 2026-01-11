package common

type GetListOrderRequest struct {
	Page     int      `form:"page"`
	Limit    int      `form:"limit"`
	MinPrice *float64 `form:"min_price"`
	MaxPrice *float64 `form:"max_price"`
	Status   *string  `form:"status"`
}
