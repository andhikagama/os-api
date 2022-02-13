package dao

import "github.com/shopspring/decimal"

type (
	Inventory struct {
		ID            uint64          `gorm:"primaryKey"`
		SkuName       string          `gorm:"not null"`
		SkuCode       string          `gorm:"not null"`
		WAC           decimal.Decimal `gorm:"not null"`
		TotalQty      decimal.Decimal `gorm:"not null"`
		AvailableQty  decimal.Decimal `gorm:"not null"`
		QuarantineQty decimal.Decimal `gorm:"not null"`

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
