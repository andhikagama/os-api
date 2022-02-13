package user

import (
	"fmt"

	"github.com/andhikagama/os-api/config/consts"
	"github.com/andhikagama/os-api/model/dto/user"
	"github.com/andhikagama/os-api/shared/utils"
)

func (h *handler) Create(ctx *utils.Context, request user.CreateRequest) (user.Response, error) {
	res, err := h.domain.User.Create(ctx, request.ToDao())
	if err != nil {
		if utils.IsDuplicateError(err) {
			return user.Response{}, fmt.Errorf("user %w", consts.ErrAlreadyExist)
		}

		return user.Response{}, err
	}

	return user.NewResponse(res), nil
}
