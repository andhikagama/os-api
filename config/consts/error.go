package consts

import "errors"

type (
	Err struct {
		Code  int
		Error error
	}

	ErrList []Err
)

var (
	ErrInternalServerError = errors.New("internal server error")
)
