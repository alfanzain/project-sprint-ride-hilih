package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alfanzain/project-sprint-halo-suster/app/consts"
	handlerContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/handlers"
	serviceContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/services"
	"github.com/alfanzain/project-sprint-halo-suster/app/entities"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/helpers"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/repositories"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/services"
	"github.com/labstack/echo/v4"
)

type MedicalPatientHandler struct {
	patientService serviceContracts.IPatientService
}

func NewMedicalPatientHandler(s serviceContracts.IPatientService) handlerContracts.IMedicalPatientHandler {
	return &MedicalPatientHandler{
		patientService: services.NewPatientService(
			repositories.NewPatientRepository(),
		),
	}
}

func (h *MedicalPatientHandler) Get(c echo.Context) (e error) {
	return c.JSON(http.StatusNotFound, SuccessResponse{
		Message: "endpoint in progress",
		Data:    nil,
	})
}

func (h *MedicalPatientHandler) Store(c echo.Context) (e error) {
	r := new(entities.PatientRegisterRequest)

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

	genderEnum := helpers.Slice[string]{
		consts.GENDER_MALE_ENUM,
		consts.GENDER_FEMALE_ENUM,
	}
	if !genderEnum.Includes(r.Gender) {
		fmt.Println("gender is invalid")

		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "gender is invalid",
		})
	}

	var genderID int
	if r.Gender == consts.GENDER_MALE_ENUM {
		genderID = consts.GENDER_MALE_ID
	}
	if r.Gender == consts.GENDER_FEMALE_ENUM {
		genderID = consts.GENDER_FEMALE_ID
	}

	birthDate, err := time.Parse("2006-01-02", r.BirthDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	data, err := h.patientService.Store(&entities.PatientStorePayload{
		ID:                  r.ID,
		PhoneNumber:         r.PhoneNumber,
		Name:                r.Name,
		BirthDate:           birthDate,
		GenderID:            genderID,
		IdentityCardScanImg: r.IdentityCardScanImg,
	})

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"patients_pkey\"" {
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

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Patient registered successfully",
		Data:    data,
	})
}
