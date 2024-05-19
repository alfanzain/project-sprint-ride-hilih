package v1

import (
	"github.com/labstack/echo/v4"
)

type V1Routes struct {
	Echo *echo.Group
}

type iV1Routes interface {
	MountUser()
	MountMedical()
}

func New(v1Routes *V1Routes) iV1Routes {
	return v1Routes
}
