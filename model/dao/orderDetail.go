package dao

type (
	OrderDetail struct {
		OrderID     uint64  `gorm:"primaryKey"`
		InventoryID uint64  `gorm:"primaryKey"`
		Qty         float64 `gorm:"not null"`
		WAC         float64 `gorm:"not null"`
		Amount      float64 `gorm:"not null"`

		Inventory Inventory
		GormModel
		BaseModelSoftDelete
	}

	OrderDetails []OrderDetail
)

// TableName .
func (OrderDetail) TableName() string {
	return "order_details"
}

// ModelName .
func (OrderDetail) ModelName() string {
	return "OrderDetail"
}
