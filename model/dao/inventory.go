package dao

type (
	Inventory struct {
		ID            uint64  `gorm:"primaryKey"`
		SkuName       string  `gorm:"not null"`
		SkuCode       string  `gorm:"not null"`
		WAC           float64 `gorm:"not null"`
		TotalQty      float64 `gorm:"not null"`
		AvailableQty  float64 `gorm:"not null"`
		QuarantineQty float64 `gorm:"not null"`

		GormModel
		BaseModelSoftDelete
	}

	Inventories []Inventory
)

// TableName .
func (Inventory) TableName() string {
	return "inventories"
}

// ModelName .
func (Inventory) ModelName() string {
	return "Inventory"
}
