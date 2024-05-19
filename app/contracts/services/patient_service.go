package services

import "github.com/alfanzain/project-sprint-halo-suster/app/entities"

type IPatientService interface {
	Store(*entities.PatientStorePayload) (*entities.Patient, error)
}
