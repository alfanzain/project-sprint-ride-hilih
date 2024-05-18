package handlers

import (
	"errors"
	"net/http"
	"strconv"

	handlerContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/handlers"
	serviceContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/services"
	"github.com/alfanzain/project-sprint-halo-suster/app/entities"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/errs"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/repositories"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/services"

	"github.com/labstack/echo/v4"
)

type UserITHandler struct {
	userITService serviceContracts.IUserITService
}

func NewUserITHandler(s serviceContracts.IUserITService) handlerContracts.IUserITHandler {
	return &UserITHandler{
		userITService: services.NewUserITService(
			repositories.NewUserRepository(),
		),
	}
}

func (h *UserITHandler) Register(c echo.Context) (e error) {
	r := new(entities.UserITRegisterRequest)

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

	data, err := h.userITService.Register(&entities.UserITRegisterPayload{
		NIP:      strconv.Itoa(r.NIP),
		Name:     r.Name,
		Password: r.Password,
	})

	if err != nil && errors.Is(err, errs.ErrInvalidNIP) {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err == errs.ErrNIPAlreadyRegistered {
		return c.JSON(http.StatusConflict, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
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
	r := new(entities.UserITLoginRequest)

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

	data, e := h.userITService.Login(&entities.UserITLoginPayload{
		NIP:      strconv.Itoa(r.NIP),
		Password: r.Password,
	})

	if e != nil {
		if e == errs.ErrUserITNotFound {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: e.Error(),
			})
		} else if e == errs.ErrInvalidPassword {
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
