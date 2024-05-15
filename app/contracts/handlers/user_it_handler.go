package handlers

import (
	"github.com/labstack/echo/v4"
)

type IUserITHandler interface {
	Register(echo.Context) error
	Login(echo.Context) error
	GetUsers(echo.Context) error
}
