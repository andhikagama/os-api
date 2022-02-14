package service

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared/utils"
	"gorm.io/gorm"
)

func (s service) UpdateExpired(ctx *utils.Context, req dao.OrderDetails) error {
	for _, v := range req {
		mapUpdate := map[string]interface{}{
			"available_qty":  gorm.Expr("available_qty + ?", v.Qty),
			"quarantine_qty": gorm.Expr("quarantine_qty - ?", v.Qty),
		}

		if err := s.repository.UpdateByID(ctx, v.InventoryID, mapUpdate); err != nil {
			return err
		}
	}
	return nil
}
