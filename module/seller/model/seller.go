package sellermodel

import (
	"OpenMarket/common"
)

const EntityName = "Seller"

type Seller struct {
	common.SQLModel `json",inline"`
	UserID          int                `gorm:"column:user_id;not null" json:"-"`
	ShopName        string             `gorm:"column:shop_name; not null" json:"shop_name"`
	ShopDescription string             `gorm:"column:shop_description; not null" json:"shop_description"`
	ShopPhone       string             `gorm:"column:shop_phone; not null" json:"shop_phone"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (Seller) TableName() string { return "sellers" }

func (s *Seller) Mask() {
	if s == nil {
		return
	}
	s.GenUID(common.DbTypeSeller)
	if u := s.User; u != nil {
		u.Mask()
	}
}

type SellerCreate struct {
	UserID          int    `gorm:"column:user_id;not null" json:"user_id"`
	ShopName        string `gorm:"column:shop_name; not null" json:"shop_name"`
	ShopDescription string `gorm:"column:shop_description; not null" json:"shop_description"`
	ShopPhone       string `gorm:"column:shop_phone; not null" json:"shop_phone"`
}

func (SellerCreate) TableName() string { return Seller{}.TableName() }

type SellerUpdate struct {
	ShopName        string `gorm:"column:shop_name; not null" json:"shop_name"`
	ShopDescription string `gorm:"column:shop_description; not null" json:"shop_description"`
	ShopPhone       string `gorm:"column:shop_phone; not null" json:"shop_phone"`
}

func (SellerUpdate) TableName() string { return Seller{}.TableName() }
