package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	// userService serviceContracts.IUserService
}

// func NewUserHandler(s serviceContracts.IUserService) handlerContracts.IUserHandler {
// 	return &UserHandler{
// 		userService: services.NewUserService(
// 			repositories.NewUserRepository(),
// 		),
// 	}
// }

func (h *UserHandler) Get(c echo.Context) (e error) {
	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "success",
		Data:    nil,
	})
}
