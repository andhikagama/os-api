package inventory

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/model/dto"
)

type (
	Response struct {
		ID            uint64  `json:"id"`
		SkuName       string  `json:"skuName"`
		SkuCode       string  `json:"skuCode"`
		WAC           float64 `json:"wac"`
		TotalQty      float64 `json:"totalQty"`
		AvailableQty  float64 `json:"availableQty"`
		QuarantineQty float64 `json:"quarantineQty"`

		dto.BaseModelSoftDelete
	}

	Responses []Response
)

func NewResponse(inv dao.Inventory) Response {
	resp := Response{
		ID:            inv.ID,
		SkuName:       inv.SkuName,
		SkuCode:       inv.SkuCode,
		WAC:           inv.WAC,
		TotalQty:      inv.TotalQty,
		AvailableQty:  inv.AvailableQty,
		QuarantineQty: inv.QuarantineQty,
	}

	resp.BaseModelSoftDelete = dto.NewBaseModelSoftDelete(inv.BaseModelSoftDelete)
	return resp
}

func NewResponses(invs dao.Inventories) Responses {
	resp := make(Responses, len(invs))

	for k, v := range invs {
		resp[k] = NewResponse(v)
	}

	return resp
}
