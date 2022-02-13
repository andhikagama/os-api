package middleware

import (
	"net/http"
	"strings"

	"github.com/andhikagama/os-api/config/consts"
	"github.com/andhikagama/os-api/shared/utils"
	"github.com/labstack/echo/v4"
)

// ValidatePrivilege .
func (m *Middleware) ValidatePrivilege(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		privilege := ``
		registered := false
		authorization := c.Request().Header.Get(consts.HeaderAuthorization)
		lang := c.Request().Header.Get(consts.HeaderAppLanguage)
		idxLang := 0
		if lang == `id` {
			idxLang = 1
		}

		c.Set(`appLanguage`, idxLang)
		for _, v := range m.Echo.Routes() {
			if c.Path() == v.Path && c.Request().Method == v.Method {
				switch v.Name {
				case consts.PrivilegeGranted, consts.PrivilegePublic, consts.PrivilegeTrusted:
					privilege = v.Name
					registered = true
				default:
					registered = false
				}
				break
			}
		}

		if !registered {
			return next(c)
		}

		if privilege == consts.PrivilegePublic {
			return next(c)
		}

		if privilege == consts.PrivilegeTrusted && authorization == `` {
			return next(c)
		}

		token := strings.Replace(authorization, `Bearer `, ``, 1)
		claims, err := m.validateToken(token)
		if err == nil {
			c.Set("claims", claims)
			return next(c)
		}

		ctx := utils.NewContext(c, nil)
		return ctx.ResponseErrorJSON(http.StatusUnauthorized, consts.ErrUnauthorized)
	}
}
