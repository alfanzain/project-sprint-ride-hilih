package v1

import (
	"github.com/alfanzain/project-sprint-halo-suster/app/http/handlers"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/middlewares"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/repositories"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/services"
)

func (i *V1Routes) MountMedical() {
	gMedical := i.Echo.Group("/medical")

	patientRepo := repositories.NewPatientRepository()
	patientService := services.NewPatientService(patientRepo)
	medicalPatientHandler := handlers.NewMedicalPatientHandler(patientService)

	gMedicalPatient := gMedical.Group("/patient")
	gMedicalPatient.POST("", medicalPatientHandler.Store, middlewares.Authorized(), middlewares.IsRoleITOrNurse())
}
