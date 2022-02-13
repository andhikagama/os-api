package user

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/model/dto"
)

type (
	Response struct {
		ID        uint64  `json:"id"`
		Phone     string  `json:"phone"`
		Email     *string `json:"email"`
		FirstName string  `json:"firstName"`
		LastName  string  `json:"lastName"`

		dto.BaseModelSoftDelete
	}
)

func NewResponse(usr dao.User) Response {
	resp := Response{
		ID:        usr.ID,
		Phone:     usr.Phone,
		Email:     usr.Email,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
	}

	resp.BaseModelSoftDelete = dto.NewBaseModelSoftDelete(usr.BaseModelSoftDelete)
	return resp
}
