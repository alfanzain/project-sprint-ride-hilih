package services

import (
	repositoryContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/repositories"
	serviceContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/services"
	"github.com/alfanzain/project-sprint-halo-suster/app/entities"
)

type PatientService struct {
	patientRepository repositoryContracts.IPatientRepository
}

func NewPatientService(
	patientRepository repositoryContracts.IPatientRepository,
) serviceContracts.IPatientService {
	return &PatientService{
		patientRepository: patientRepository,
	}
}

func (s *PatientService) Store(p *entities.PatientStorePayload) (*entities.Patient, error) {
	patient, err := s.patientRepository.Store(&entities.PatientStorePayload{
		ID:                  p.ID,
		PhoneNumber:         p.PhoneNumber,
		Name:                p.Name,
		BirthDate:           p.BirthDate,
		GenderID:            p.GenderID,
		IdentityCardScanImg: p.IdentityCardScanImg,
	})

	if err != nil {
		return nil, err
	}

	return &entities.Patient{
		ID:                  patient.ID,
		PhoneNumber:         patient.PhoneNumber,
		Name:                patient.Name,
		BirthDate:           patient.BirthDate,
		GenderID:            patient.GenderID,
		IdentityCardScanImg: patient.IdentityCardScanImg,
		CreatedAt:           patient.CreatedAt,
	}, nil
}
