package inventory

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/model/dto"
	"github.com/shopspring/decimal"
)

type (
	Response struct {
		ID            uint64          `json:"id"`
		SkuName       string          `json:"skuName"`
		SkuCode       string          `json:"skuCode"`
		WAC           decimal.Decimal `json:"wac"`
		TotalQty      decimal.Decimal `json:"totalQty"`
		AvailableQty  decimal.Decimal `json:"availableQty"`
		QuarantineQty decimal.Decimal `json:"quarantineQty"`

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
