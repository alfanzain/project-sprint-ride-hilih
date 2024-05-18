package handlers

import (
	"github.com/labstack/echo/v4"
)

type IUserNurseHandler interface {
	Register(echo.Context) error
	Login(echo.Context) error
	Update(echo.Context) error
	Destroy(echo.Context) error
	GrantAccess(echo.Context) error
}
