package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alfanzain/project-sprint-halo-suster/app/consts"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

func IsRoleIT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userCtx := new(UserValidator)
			mapstructure.Decode(c.Get("user"), &userCtx)

			fmt.Println(userCtx.UserID)
			fmt.Println(userCtx.RoleID)

			if userCtx.RoleID != strconv.Itoa(consts.NIP_CODE_ROLE_IT) {
				return c.JSON(http.StatusUnauthorized, ErrorResponse{
					Status:  false,
					Message: "Unauthorized",
				})
			}

			return next(c)
		}
	}
}
