package user

import (
	"github.com/andhikagama/os-api/model/dao"
)

type (
	LoginResponse struct {
		Response
		Token string `json:"token"`
	}
)

func NewLoginResponse(usr dao.User) LoginResponse {
	return LoginResponse{
		Response: NewResponse(usr),
		Token:    *usr.Token,
	}
}
