package sellermodel

type Seller struct {
	ID              int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID          int    `gorm:"column:user_id;not null" json:"user_id"`
	ShopName        string `gorm:"column:shop_name; not null" json:"shop_name"`
	ShopDescription string `gorm:"column:shop_description; not null" json:"shop_description"`
	ShopPhone       string `gorm:"column:shop_phone; not null" json:"shop_phone"`
	Status          string `gorm:"column:status; not null" json:"status"`
	CreatedAt       int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (Seller) TableName() string { return "sellers" }

type SellerCreate struct {
	ID              int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID          int    `gorm:"column:user_id;not null" json:"user_id"`
	ShopName        string `gorm:"column:shop_name; not null" json:"shop_name"`
	ShopDescription string `gorm:"column:shop_description; not null" json:"shop_description"`
	ShopPhone       string `gorm:"column:shop_phone; not null" json:"shop_phone"`
	Status          string `gorm:"column:status; not null" json:"status"`
}

func (SellerCreate) TableName() string { return Seller{}.TableName() }
