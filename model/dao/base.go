package dao

import (
	"time"

	"github.com/andhikagama/os-api/model/dto"

	"gorm.io/gorm"
)

type (
	BaseModelSoftDelete struct {
		CreatedAt time.Time  `gorm:"not null"`
		UpdatedAt time.Time  `gorm:"not null"`
		DeletedAt *time.Time `gorm:"null"`
	}

	GormModel struct {
		DeletedAt gorm.DeletedAt
	}

	BasePagination struct {
		Page      int64
		Limit     int64
		Sort      []string
		NextPage  int64
		TotalPage int64
		Count     int64
	}
)

func (bm BaseModelSoftDelete) ToDTO() dto.BaseModelSoftDelete {
	return dto.BaseModelSoftDelete{
		CreatedAt: bm.CreatedAt,
		UpdatedAt: bm.UpdatedAt,
		DeletedAt: bm.DeletedAt,
	}
}

func NewBasePaginationFromDTO(pg dto.BasePagination) BasePagination {
	return BasePagination{
		Page:  pg.Page,
		Limit: pg.Limit,
		Sort:  pg.Sort,
	}
}
