package dao

import (
	"time"

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
