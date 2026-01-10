package cartmodel

type Cart struct {
	Id     int `json:"-" gorm:"column:id;"`
	UserId int `json:"-" gorm:"not null;index"`
}

func (Cart) TableName() string { return "carts" }

type CartCreate struct {
	UserId int `json:"-" gorm:"not null;index"`
}

func (CartCreate) TableName() string { return Cart{}.TableName() }

type CartItem struct {
	Id        int `json:"-" gorm:"column:id;"`
	CartId    int `json:"-" gorm:"column:cart_id;not null;index"`
	VariantId int `json:"variant_id" gorm:"not null;index"`
	Quantity  int `json:"quantity" gorm:"column:quantity;not null;"`
}

func (CartItem) TableName() string { return "cart_items" }

type CartItemCreate struct {
	CartId    int `json:"-" gorm:"column:cart_id;not null;index"`
	VariantId int `json:"variant_id" gorm:"not null;index"`
	Quantity  int `json:"quantity" gorm:"column:quantity;not null;"`
}

func (CartItemCreate) TableName() string { return CartItem{}.TableName() }

type CartItemUpdate struct {
	Quantity int `json:"quantity" gorm:"column:quantity;"`
}

func (CartItemUpdate) TableName() string { return CartItem{}.TableName() }

type CartItemView struct {
	VariantId int `json:"variant_id"`
	Quantity  int `json:"quantity"`
}

type CartView struct {
	Id    int            `json:"id"`
	Items []CartItemView `json:"items"`
}
