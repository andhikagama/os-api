package middleware

import (
	"github.com/andhikagama/os-api/config/consts"
	"github.com/andhikagama/os-api/shared/utils"
)

func (m *Middleware) validateToken(token string) (*utils.Token, error) {
	claims, err := utils.ClaimToken(token)
	if err != nil {
		return &utils.Token{}, consts.ErrUnauthorized
	}

	return claims, nil
}
