package dao

import (
	"github.com/andhikagama/os-api/config/consts/orderEnum"
)

type (
	Order struct {
		ID            uint64                  `gorm:"primaryKey"`
		UserID        uint64                  `gorm:"not null"`
		TotalQty      float64                 `gorm:"not null"`
		TotalAmount   float64                 `gorm:"not null"`
		PaymentType   orderEnum.PaymentType   `gorm:"not null"`
		PaymentStatus orderEnum.PaymentStatus `gorm:"not null"`

		User         User
		OrderDetails OrderDetails

		GormModel
		BaseModelSoftDelete
	}

	Orders []Order
)

// TableName .
func (Order) TableName() string {
	return "orders"
}

// ModelName .
func (Order) ModelName() string {
	return "Order"
}
