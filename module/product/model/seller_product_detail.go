package productmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

// SellerProductDetail is the response for GET /v1/seller/products/:id
type SellerProductDetail struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	BasePrice   decimal.Decimal `json:"base_price"`
	Status      int             `json:"status"`
	CategoryID  *int            `json:"category_id,omitempty"`
	SellerID    string          `json:"seller_id"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
	Images      []GalleryItem   `json:"images"`
}

// GalleryItem represents an image in the gallery
type GalleryItem struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	IsMain bool   `json:"is_main"`
}
