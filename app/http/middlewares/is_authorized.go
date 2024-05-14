package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"

	"github.com/alfanzain/project-sprint-halo-suster/app/helpers"
)

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type roleValidator struct {
	Role string `mapstructure:"role" validate:"required"`
}

func Authorized() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)

			if token == "" {
				return c.JSON(http.StatusUnauthorized, ErrorResponse{
					Status:  false,
					Message: "Unauthorized",
				})
			}

			claims, err := helpers.ValidateJWT(&helpers.ParamsValidateJWT{
				Token:     token,
				SecretKey: os.Getenv("JWT_SECRET"),
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, ErrorResponse{
					Status:  false,
					Message: "Unauthorized",
				})
			}

			user := make(map[string]interface{})
			userRole := new(roleValidator)
			mapstructure.Decode(claims, &user)
			mapstructure.Decode(claims, &userRole)

			c.Set("user", user)

			return next(c)
		}
	}
}
