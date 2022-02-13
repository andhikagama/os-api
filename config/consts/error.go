package consts

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type (
	Err struct {
		Code  int
		Error error
	}

	ErrList []Err
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrEmptyRequestBody    = errors.New(`request body can't be empty`)
	ErrUnauthorized        = errors.New(`unauthorized`)
	ErrAlreadyExist        = errors.New("already exist")

	// User Error List
	ErrListUser = ErrList{
		{Code: http.StatusNotFound, Error: gorm.ErrRecordNotFound},
		{Code: http.StatusConflict, Error: ErrAlreadyExist},
	}
)
