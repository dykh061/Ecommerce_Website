package ordermodel

import "time"

type OrderAddress struct {
	Id        int        `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	OrderId   int        `json:"order_id" gorm:"column:order_id;index"`
	Address   string     `json:"address" gorm:"column:address;type:text"`
	Phone     string     `json:"phone" gorm:"column:phone;size:15"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (OrderAddress) TableName() string { return "order_addresses" }

type OrderAddressCreate struct {
	OrderId  int    `json:"-" gorm:"column:order_id;index"`
	FullName string `json:"full_name" gorm:"column:full_name"`
	Address  string `json:"address" gorm:"column:address;type:text"`
	City     string `json:"city" gorm:"type:varchar(100);not null"`
	Phone    string `json:"phone" gorm:"column:phone;size:15"`
}

func (OrderAddressCreate) TableName() string { return OrderAddress{}.TableName() }

type OrderAddressUpdate struct {
	Address  *string `json:"address" gorm:"column:address;type:text"`
	FullName *string `json:"full_name" gorm:"column:full_name;type:text"`
	City     *string `json:"city" gorm:"type:varchar(100);not null"`
	Phone    *string `json:"phone" gorm:"column:phone;size:15"`
}

func (OrderAddressUpdate) TableName() string { return OrderAddress{}.TableName() }
