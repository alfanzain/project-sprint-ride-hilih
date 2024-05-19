package repositories

import "github.com/alfanzain/project-sprint-halo-suster/app/entities"

type IPatientRepository interface {
	Store(*entities.PatientStorePayload) (*entities.Patient, error)
}
