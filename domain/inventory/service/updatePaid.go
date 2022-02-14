package service

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared/utils"
	"gorm.io/gorm"
)

func (s service) UpdatePaid(ctx *utils.Context, req dao.OrderDetails) error {
	for _, v := range req {
		mapUpdate := map[string]interface{}{
			"total_qty":      gorm.Expr("available_qty - ?", v.Qty),
			"quarantine_qty": gorm.Expr("quarantine_qty - ?", v.Qty),
		}

		if err := s.repository.UpdateByID(ctx, v.InventoryID, mapUpdate); err != nil {
			return err
		}
	}
	return nil
}

// if paid decrease inventory total qty, and quarantine qty
// if expired decrease inventory quarantine qty and increase available qty
