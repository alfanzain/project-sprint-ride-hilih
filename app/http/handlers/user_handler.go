package handlers

import (
	"net/http"
	"strconv"

	handlerContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/handlers"
	serviceContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/services"
	"github.com/alfanzain/project-sprint-halo-suster/app/entities"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/repositories"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/services"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService serviceContracts.IUserService
}

func NewUserHandler(s serviceContracts.IUserService) handlerContracts.IUserHandler {
	return &UserHandler{
		userService: services.NewUserService(
			repositories.NewUserRepository(),
		),
	}
}

func (h *UserHandler) Get(c echo.Context) (e error) {
	filters := &entities.UserGetFilterParams{}

	if id := c.QueryParam("id"); id != "" {
		filters.ID = id
	}
	if limitStr := c.QueryParam("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'limit'",
			})
		}
		filters.Limit = limit
	}
	if offsetStr := c.QueryParam("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'offset'",
			})
		}
		filters.Offset = offset
	}
	if name := c.QueryParam("name"); name != "" {
		filters.Name = name
	}
	if nip := c.QueryParam("nip"); nip != "" {
		filters.NIP = nip
	}
	if role := c.QueryParam("role"); role != "" {
		filters.Role = role
	}
	if createdAt := c.QueryParam("createdAt"); createdAt != "" {
		if createdAt == "asc" || createdAt == "desc" {
			filters.CreatedAt = createdAt
		}
	}

	data, e := h.userService.GetUsers(filters)
	if e != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: e.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "success",
		Data:    data,
	})
}
