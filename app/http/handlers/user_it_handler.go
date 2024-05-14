package handlers

import (
	"net/http"
	"strconv"

	handlerContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/handlers"
	handlerServices "github.com/alfanzain/project-sprint-halo-suster/app/contracts/services"
	"github.com/alfanzain/project-sprint-halo-suster/app/entities"
	"github.com/alfanzain/project-sprint-halo-suster/app/repositories"
	"github.com/alfanzain/project-sprint-halo-suster/app/services"

	"github.com/labstack/echo/v4"
)

type UserITHandler struct {
	UserITService handlerServices.IUserITService
}

func NewUserITHandler(s handlerServices.IUserITService) handlerContracts.IUserITHandler {
	return &UserITHandler{
		UserITService: services.NewUserITService(
			repositories.NewUserITRepository(),
		),
	}
}

type (
	RegisterRequest struct {
		NIP      int    `json:"nip" validate:"required,min=13,max=13"`
		Name     string `json:"name" validate:"required,min=5,max=50"`
		Password string `json:"password" validate:"required,min=5,max=15"`
	}

	LoginRequest struct {
		NIP      int    `json:"nip" validate:"required,min=13,max=13"`
		Password string `json:"password" validate:"required,min=5,max=15"`
	}
)

func (h *UserITHandler) Register(c echo.Context) (e error) {
	r := new(RegisterRequest)

	if e = c.Bind(r); e != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: e.Error(),
		})
	}

	if e = c.Validate(r); e != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: e.Error(),
		})
	}

	data, err := h.UserITService.Register(&entities.UserITRegisterPayload{
		NIP:      strconv.Itoa(r.NIP),
		Name:     r.Name,
		Password: r.Password,
	})

	if err != nil {
		return c.JSON(http.StatusConflict, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "User registered successfully",
		Data:    data,
	})
}

func (h *UserITHandler) Login(c echo.Context) (e error) {
	r := new(LoginRequest)

	if e = c.Bind(r); e != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: e.Error(),
		})
	}

	if e = c.Validate(r); e != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: e.Error(),
		})
	}

	data, e := h.UserITService.Login(&entities.UserITLoginPayload{
		NIP:      strconv.Itoa(r.NIP),
		Password: r.Password,
	})

	if e != nil {
		if e == services.ErrUserITNotFound {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: e.Error(),
			})
		} else if e == services.ErrInvalidPassword {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: e.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "User logged successfully",
		Data:    data,
	})
}
