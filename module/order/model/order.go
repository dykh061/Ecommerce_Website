package ordermodel

import (
	"time"

	"github.com/shopspring/decimal"
)

const (
	OrderPending   = "pending"
	OrderConfirmed = "confirmed"
	OrderShipping  = "shipping"
	OrderCompleted = "completed"
	OrderCancelled = "cancelled"
)

type Order struct {
	Id          int             `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserId      int             `json:"user_id" gorm:"column:user_id;index"`
	TotalAmount decimal.Decimal `json:"total_amount" gorm:"column:total_amount"`
	Status      string          `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time      `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (Order) TableName() string { return "orders" }

type OrderCreate struct {
	UserId      int             `json:"-" gorm:"column:user_id"`
	TotalAmount decimal.Decimal `json:"total_amount" gorm:"column:total_amount"`
	Status      string          `json:"status" gorm:"column:status"`
}

func (OrderCreate) TableName() string { return Order{}.TableName() }

type OrderTotalAmount struct {
	TotalAmount decimal.Decimal `json:"total_amount" gorm:"column:total_amount"`
}

type OrderStatusUpdate struct {
	Status string `json:"status"`
}

func (OrderStatusUpdate) TableName() string { return Order{}.TableName() }

type OrderCancel struct {
	Reason string `json:"reason"`
}

func (OrderCancel) TableName() string { return Order{}.TableName() }

type FilterOrder struct {
	UserId   *int             `json:"-" gorm:"column:user_id"`
	MinPrice *decimal.Decimal `json:"min_price" form:"min_price"`
	MaxPrice *decimal.Decimal `json:"max_price" form:"max_price"`
	Status   *string          `json:"status" gorm:"column:status"`
}
