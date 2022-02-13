package inventory

import (
	"github.com/andhikagama/os-api/model/dto"
	"github.com/andhikagama/os-api/shared/utils"
	"gorm.io/gorm"
)

type (
	PaginatedRequest struct {
		dto.BasePagination
		Filter FilterRequest `json:"filter,omitempty"`
	}

	FilterRequest struct {
		SkuName string `query:"skuName" json:"skuName,omitempty"`
		SkuCode string `query:"skuCode" json:"skuCode,omitempty"`
	}
)

func NewPaginatedRequest(ctx *utils.Context) (PaginatedRequest, error) {
	request := PaginatedRequest{}

	if err := ctx.Bind(&request); err != nil {
		return PaginatedRequest{}, err
	}

	request.BasePagination.SetDefault()
	return request, nil
}

func (PaginatedRequest) AvailableSort() map[string]string {
	return map[string]string{
		"updated_at":  "inventories.updated_at",
		"-updated_at": "-inventories.updated_at",
		"created_at":  "inventories.created_at",
		"-created_at": "-inventories.created_at",
	}
}

func (f FilterRequest) Apply(tx *gorm.DB) {
	if f.SkuName != "" {
		tx.Where("sku_name REGEXP ?", f.SkuName)
	}

	if f.SkuCode != "" {
		tx.Where("sku_code REGEXP ?", f.SkuCode)
	}
}
