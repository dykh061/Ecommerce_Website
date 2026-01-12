package adminmodel

import (
	"time"

	"OpenMarket/common"
)

// SellerAdminView represents seller data for admin panel
type SellerAdminView struct {
	Id              int        `json:"-" gorm:"column:id"`
	FakeId          *common.UID `json:"id" gorm:"-"`
	UserId          int        `json:"-" gorm:"column:user_id"`
	ShopName        string     `json:"shop_name" gorm:"column:shop_name"`
	ShopDescription string     `json:"shop_description" gorm:"column:shop_description"`
	ShopPhone       string     `json:"shop_phone" gorm:"column:shop_phone"`
	Status          int        `json:"status" gorm:"column:status"`
	CreatedAt       *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	// User info
	UserName  string `json:"user_name" gorm:"column:user_name"`
	UserEmail string `json:"user_email" gorm:"column:user_email"`
	UserPhone string `json:"user_phone" gorm:"column:user_phone"`
}

func (s *SellerAdminView) Mask() {
	s.FakeId = common.NewUID(uint32(s.Id), common.DbTypeSeller, 1)
}

// SellerStatusUpdate for updating seller status
type SellerStatusUpdate struct {
	Status int    `json:"status" binding:"required,oneof=0 1"`
	Reason string `json:"reason"`
}

// SellerFilter for filtering sellers in admin panel
type SellerFilter struct {
	Status   *int    `json:"status" form:"status"`
	Keyword  *string `json:"keyword" form:"keyword"`
	SortBy   string  `json:"sort_by" form:"sort_by"`
	SortDesc bool    `json:"sort_desc" form:"sort_desc"`
}

// Seller status constants
const (
	SellerStatusInactive = 0
	SellerStatusActive   = 1
)
