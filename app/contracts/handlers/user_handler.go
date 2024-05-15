package handlers

import (
	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	Get(echo.Context) error
}
