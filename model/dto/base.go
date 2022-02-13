package dto

import (
	"math"
	"strings"
	"time"

	"github.com/andhikagama/os-api/model/dao"

	"gorm.io/gorm"
)

type (
	BaseModelSoftDelete struct {
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at,omitempty"`
		DeletedAt *time.Time `json:"deleted_at,omitempty"`
	}

	BasePagination struct {
		Page      int64    `query:"page" json:"page" validate:"number"`
		Limit     int64    `query:"limit" json:"limit" validate:"number"`
		Sort      []string `query:"sort" json:"sort"`
		NextPage  int64    `json:"next_page"`
		TotalPage int64    `json:"total_page"`
		Count     int64    `json:"count"`
	}

	SuccessMessage struct {
		Message string `json:"message"`
	}
)

// ClockType .
func (b *BaseModelSoftDelete) ClockType() {
	ca := b.CreatedAt.Round(time.Second)
	b.CreatedAt = ca

	ua := b.UpdatedAt.Round(time.Second)
	b.UpdatedAt = ua
}

// SetDefault .
func (p *BasePagination) SetDefault() *BasePagination {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.Limit == 0 {
		p.Limit = 10
	}

	if len(p.Sort) == 0 {
		p.Sort = []string{`-created_at`}
	} else {
		sorts := strings.Split(p.Sort[0], `,`)
		p.Sort = sorts
	}

	return p
}

func (pg *BasePagination) ToPaginatedQuery(q *gorm.DB, availableSort map[string]string) {
	nextPage := int64(0)
	count := int64(0)

	q.Count(&count)

	offset := (pg.Page - 1) * pg.Limit
	order := buildSort(pg.Sort, availableSort)
	q.Limit(int(pg.Limit)).Offset(int(offset)).Order(order)

	totalPages := int64(math.Ceil(float64(count) / float64(pg.Limit)))
	if totalPages == 0 {
		totalPages = 1
	}

	if pg.Page < totalPages {
		nextPage = pg.Page + 1
	}

	pg.TotalPage = totalPages
	pg.Count = count
	pg.NextPage = nextPage
}

func buildSort(sort []string, availableSort map[string]string) string {
	order := make([]string, len(sort))

	i := 0
	for _, v := range sort {
		temp := ``
		val, _ := availableSort[v]

		temp = val
		if strings.Contains(temp, `-`) {
			order[i] = strings.Replace(temp, `-`, ``, -1) + ` desc`
		} else {
			order[i] = temp + ` asc`
		}

		i++
	}

	return strings.Join(order, `, `)
}

func (bm BaseModelSoftDelete) ToDao() dao.BaseModelSoftDelete {
	return dao.BaseModelSoftDelete{
		CreatedAt: bm.CreatedAt,
		UpdatedAt: bm.UpdatedAt,
		DeletedAt: bm.DeletedAt,
	}
}

func NewBaseModelSoftDelete(bm dao.BaseModelSoftDelete) BaseModelSoftDelete {
	return BaseModelSoftDelete{
		CreatedAt: bm.CreatedAt,
		UpdatedAt: bm.UpdatedAt,
		DeletedAt: bm.DeletedAt,
	}
}

func (pg BasePagination) ToDao() dao.BasePagination {
	return dao.BasePagination{
		Page:  pg.Page,
		Limit: pg.Limit,
		Sort:  pg.Sort,
	}
}
