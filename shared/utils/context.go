package utils

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/andhikagama/os-api/config/consts"
	"github.com/andhikagama/os-api/shared/resource"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

type (
	Context struct {
		EchoCtx   echo.Context
		Tx        *gorm.DB
		Validator *resource.Validator
		Ctx       context.Context
	}

	errorWrapper struct {
		Message  string   `json:"message"`
		Messages []string `json:"messages"`
	}

	responseError struct {
		Error errorWrapper `json:"error"`
	}

	responseSuccess struct {
		Data interface{} `json:"data"`
	}
)

func NewContext(ec echo.Context, val *resource.Validator) *Context {
	ctx := Context{
		EchoCtx:   ec,
		Validator: val,
		Ctx:       context.TODO(),
	}

	return &ctx
}

func (c *Context) Validate(i interface{}) error {
	err := c.EchoCtx.Validate(i)
	if err != nil {
		return err
	}

	return nil
}

func (c *Context) Bind(i interface{}) error {
	err := c.EchoCtx.Bind(i)
	if err != nil {
		return err
	}

	return nil
}

func (c *Context) SetORMTransaction(tx *gorm.DB) {
	c.Tx = tx.Begin()
}

func (c *Context) GetORMTransaction(tx *gorm.DB) *gorm.DB {
	if c.Tx == nil {
		return tx
	}
	return c.Tx
}

func (c *Context) RollBackORMTransaction() {
	c.Tx.Rollback()
}

func (c *Context) CommitORMTransaction() {
	c.Tx.Commit()
	c.Tx = nil
}

func (c Context) GetUserPhone() string {
	claims := c.EchoCtx.Get("claims").(*Token)
	return claims.UserPhone
}

func (c Context) GetUserID() uint64 {
	claims := c.EchoCtx.Get("claims").(*Token)
	return claims.UserID
}

func (c *Context) GetUserToken() string {
	token := strings.Split(c.EchoCtx.Request().Header.Get(consts.HeaderAuthorization), ` `)
	return token[1]
}

func (c Context) GetAppLanguage() int {
	return c.EchoCtx.Get(`appLanguage`).(int)
}

func (c *Context) ResponseSuccessJSON(data interface{}) error {
	return c.EchoCtx.JSON(http.StatusOK, responseSuccess{Data: data})
}

func (c *Context) ResponseCreatedJSON(data interface{}) error {
	return c.EchoCtx.JSON(http.StatusCreated, responseSuccess{Data: data})
}

func (c *Context) ResponseErrorJSON(code int, err error) error {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return c.EchoCtx.JSON(code, responseError{
			Error: errorWrapper{
				Message:  err.Error(),
				Messages: []string{err.Error()},
			},
		})
	}

	mapError := errs.Translate(*c.Validator.Translators[c.GetAppLanguage()])
	errDescs := make([]string, 0)
	for _, v := range mapError {
		errDescs = append(errDescs, strings.ToLower(v))
	}

	return c.EchoCtx.JSON(http.StatusBadRequest, responseError{
		Error: errorWrapper{
			Message:  err.Error(),
			Messages: errDescs,
		},
	})
}

func (c *Context) ResponseErrorJSONFromList(err error, errList consts.ErrList) error {
	var (
		code    = http.StatusInternalServerError
		respErr = consts.ErrInternalServerError
	)

	for _, v := range errList {
		if errors.Is(err, v.Error) {
			code = v.Code
			respErr = err
		}
	}

	return c.EchoCtx.JSON(code, responseError{
		Error: errorWrapper{
			Message:  respErr.Error(),
			Messages: []string{respErr.Error()},
		},
	})
}
