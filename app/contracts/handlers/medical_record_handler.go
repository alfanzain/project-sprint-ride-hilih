package handlers

import (
	"github.com/labstack/echo/v4"
)

type IMedicalRecordHandler interface {
	Get(echo.Context) error
	Store(echo.Context) error
}
