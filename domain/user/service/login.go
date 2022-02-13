package service

import (
	"time"

	"github.com/andhikagama/os-api/config/consts"
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/model/dto/user"
	"github.com/andhikagama/os-api/shared/utils"
)

func (s service) Login(ctx *utils.Context, request user.LoginRequest) (dao.User, error) {
	md5Password := utils.NewMD5Hash(request.Password)
	encryptedPassword, err := utils.NewEncryptedString(md5Password)
	if err != nil {
		return dao.User{}, consts.ErrBadRequest
	}

	result, err := s.repository.GetByPhonePassword(ctx, request.Phone, encryptedPassword)
	if err != nil {
		return dao.User{}, consts.ErrUserInvalidPhoneOrPassword
	}

	token := utils.Token{
		UserID:    result.ID,
		UserPhone: result.Phone,
	}

	token.IssuedAt = time.Now().Unix()
	token.ExpiresAt = time.Now().AddDate(0, 1, 0).Unix()

	tokenString, err := utils.GenerateToken(token)
	if err != nil {
		return dao.User{}, err
	}

	result.Token = &tokenString
	updateRequest := map[string]interface{}{
		"token": tokenString,
	}

	ctx.SetORMTransaction(s.resource.DB)
	if err := s.repository.UpdateByID(ctx, result.ID, updateRequest); err != nil {
		ctx.RollBackORMTransaction()
		return dao.User{}, err
	}

	ctx.CommitORMTransaction()

	return result, nil
}
