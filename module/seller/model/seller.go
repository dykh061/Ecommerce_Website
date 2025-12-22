package sellermodel

import (
	"OpenMarket/common"
)

type Seller struct {
	common.SQLModel `json",inline"`
	UserID          int    `gorm:"column:user_id;not null" json:"user_id"`
	ShopName        string `gorm:"column:shop_name; not null" json:"shop_name"`
	ShopDescription string `gorm:"column:shop_description; not null" json:"shop_description"`
	ShopPhone       string `gorm:"column:shop_phone; not null" json:"shop_phone"`
}

func (Seller) TableName() string { return "sellers" }

type SellerCreate struct {
	UserID          int    `gorm:"column:user_id;not null" json:"user_id"`
	ShopName        string `gorm:"column:shop_name; not null" json:"shop_name"`
	ShopDescription string `gorm:"column:shop_description; not null" json:"shop_description"`
	ShopPhone       string `gorm:"column:shop_phone; not null" json:"shop_phone"`
}

func (SellerCreate) TableName() string { return Seller{}.TableName() }
