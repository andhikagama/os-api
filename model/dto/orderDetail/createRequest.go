package orderDetail

import "github.com/andhikagama/os-api/model/dao"

type (
	CreateRequest struct {
		InventoryID uint64  `json:"inventoryId" validate:"required"`
		Qty         float64 `json:"qty" validate:"required"`
	}

	BulkCreateRequest []CreateRequest
)

func (req CreateRequest) ToDao() dao.OrderDetail {
	return dao.OrderDetail{
		InventoryID: req.InventoryID,
		Qty:         req.Qty,
	}
}

func (req BulkCreateRequest) ToDao() dao.OrderDetails {
	data := make(dao.OrderDetails, len(req))

	for k, v := range req {
		data[k] = v.ToDao()
	}

	return data
}
