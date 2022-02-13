package user

import (
	"github.com/andhikagama/os-api/shared/utils"
)

type (
	LoginRequest struct {
		Phone    string `json:"phone" validate:"required,e164"`
		Password string `json:"password" validate:"required"`
	}
)

func NewLoginRequest(ctx *utils.Context) (LoginRequest, error) {
	request := LoginRequest{}

	if err := ctx.Bind(&request); err != nil {
		return LoginRequest{}, err
	}

	if err := ctx.Validate(&request); err != nil {
		return LoginRequest{}, err
	}

	return request, nil
}
