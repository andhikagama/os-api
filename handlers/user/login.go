package user

import (
	"github.com/andhikagama/os-api/model/dto/user"
	"github.com/andhikagama/os-api/shared/utils"
)

func (h *handler) Login(ctx *utils.Context, request user.LoginRequest) (user.LoginResponse, error) {
	res, err := h.domain.User.Login(ctx, request)
	if err != nil {
		return user.LoginResponse{}, err
	}

	return user.NewLoginResponse(res), nil
}
