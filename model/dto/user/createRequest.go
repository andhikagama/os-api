package user

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared/utils"
)

type (
	CreateRequest struct {
		Phone     string  `json:"phone" validate:"required,e164"`
		Password  string  `json:"password" validate:"required"`
		Email     *string `json:"email"`
		FirstName string  `json:"firstName" validate:"required"`
		LastName  string  `json:"lastName" validate:"required"`
	}
)

func NewCreateRequest(ctx *utils.Context) (CreateRequest, error) {
	request := CreateRequest{}

	if err := ctx.Bind(&request); err != nil {
		return CreateRequest{}, err
	}

	if err := ctx.Validate(&request); err != nil {
		return CreateRequest{}, err
	}

	return request, nil
}

func (req CreateRequest) ToDao() dao.User {
	return dao.User{
		Phone:     req.Phone,
		Password:  req.Password,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
}
