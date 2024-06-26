package handlers

import (
	"errors"
	"log"
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

type UserNurseHandler struct {
	userNurseService serviceContracts.IUserNurseService
}

func NewUserNurseHandler(s serviceContracts.IUserNurseService) handlerContracts.IUserNurseHandler {
	return &UserNurseHandler{
		userNurseService: services.NewUserNurseService(
			repositories.NewUserRepository(),
		),
	}
}

func (h *UserNurseHandler) Register(c echo.Context) (e error) {
	r := new(entities.UserNurseRegisterRequest)

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

	data, err := h.userNurseService.Register(&entities.UserNurseRegisterPayload{
		NIP:                 strconv.Itoa(r.NIP),
		Name:                r.Name,
		IdentityCardScanImg: r.IdentityCardScanImg,
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

func (h *UserNurseHandler) Login(c echo.Context) (e error) {
	log.Println("user nurse handler login")

	r := new(entities.UserNurseLoginRequest)

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

	data, e := h.userNurseService.Login(&entities.UserNurseLoginPayload{
		NIP:      strconv.Itoa(r.NIP),
		Password: r.Password,
	})

	if e != nil {
		if e == errs.ErrUserNotFound {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: e.Error(),
			})
		}
		if e == errs.ErrUserNurseNotFound {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: e.Error(),
			})
		}
		if e == errs.ErrInvalidPassword {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: e.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: e.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "User logged successfullys",
		Data:    data,
	})
}

func (h *UserNurseHandler) Update(c echo.Context) (e error) {
	r := new(entities.UserUpdateRequest)

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

	id := c.Param("userID")

	data, err := h.userNurseService.Update(&entities.UserUpdatePayload{
		ID:   id,
		NIP:  strconv.Itoa(r.NIP),
		Name: r.Name,
	})

	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}

		if errors.Is(err, errs.ErrInvalidNIP) {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}

		if errors.Is(err, errs.ErrUserNurseNotFound) {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}

		if errors.Is(err, errs.ErrNIPAlreadyRegistered) {
			return c.JSON(http.StatusConflict, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "User updated successfully",
		Data:    data,
	})
}

func (h *UserNurseHandler) Delete(c echo.Context) (e error) {
	id := c.Param("userID")

	data, err := h.userNurseService.Delete(id)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}

		if errors.Is(err, errs.ErrUserNurseNotFound) {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "User deleted successfully",
		Data:    data,
	})
}

func (h *UserNurseHandler) GrantAccess(c echo.Context) (e error) {
	r := new(entities.UserNurseGrantAccessRequest)

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

	id := c.Param("userID")

	data, err := h.userNurseService.UpdatePassword(&entities.UserNurseGrantAccessPayload{
		ID:       id,
		Password: r.Password,
	})
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}

		if errors.Is(err, errs.ErrUserNurseNotFound) {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "User access granted successfully",
		Data:    data,
	})
}
