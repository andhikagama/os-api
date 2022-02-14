package service

import (
	"github.com/andhikagama/os-api/shared/utils"
	"gorm.io/gorm"
)

func (s service) DecreaseAvailableQty(ctx *utils.Context, id uint64, qty float64) error {
	mapUpdate := map[string]interface{}{
		"available_qty":  gorm.Expr("available_qty - ?", qty),
		"quarantine_qty": gorm.Expr("quarantine_qty + ?", qty),
	}

	return s.repository.UpdateByID(ctx, id, mapUpdate)
}
