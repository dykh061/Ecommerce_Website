package sellermodel

type SellerFilter struct {
	Id      *int    `form:"id" json:"id"`
	Keyword *string `form:"keyword" json:"keyword"`
	//Status  *int    `form:"status" json:"status"`
}
