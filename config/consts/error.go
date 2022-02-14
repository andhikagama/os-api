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
	ErrBadRequest          = errors.New(`bad request`)

	ErrUnknownConstant = errors.New("unknown constant")
	ErrConstantParsing = errors.New("expected string for the constant")

	// User Error List
	ErrUserInvalidPhoneOrPassword = errors.New("invalid phone or password")

	ErrListUser = ErrList{
		{Code: http.StatusNotFound, Error: gorm.ErrRecordNotFound},
		{Code: http.StatusConflict, Error: ErrAlreadyExist},
		{Code: http.StatusBadRequest, Error: ErrUserInvalidPhoneOrPassword},
	}

	// Inventory Error List
	ErrListInventory = ErrList{
		{Code: http.StatusNotFound, Error: gorm.ErrRecordNotFound},
		{Code: http.StatusConflict, Error: ErrAlreadyExist},
	}

	// Order Error List
	ErrOrderInsufficientQty = errors.New("insufficient qty")
	ErrListOrder            = ErrList{
		{Code: http.StatusBadRequest, Error: ErrOrderInsufficientQty},
	}
)
