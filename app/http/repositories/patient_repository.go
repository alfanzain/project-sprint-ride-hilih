package repositories

import (
	repositoryContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/repositories"
	"github.com/alfanzain/project-sprint-halo-suster/app/databases"
	"github.com/alfanzain/project-sprint-halo-suster/app/entities"

	"database/sql"
	"log"
)

type PatientRepository struct {
	DB *sql.DB
}

func NewPatientRepository() repositoryContracts.IPatientRepository {
	return &PatientRepository{DB: databases.PostgreSQLInstance}
}

func (r *PatientRepository) Store(p *entities.PatientStorePayload) (*entities.Patient, error) {
	var id string
	err := r.DB.QueryRow(`
			INSERT INTO patients (
				id,
				phone_number,
				name,
				birth_date,
				gender_id,
				identity_card_scan_img
			) VALUES (
				$1, $2, $3, $4, $5, $6
			) RETURNING id
		`,
		p.ID,
		p.PhoneNumber,
		p.Name,
		p.BirthDate,
		p.GenderID,
		p.IdentityCardScanImg,
	).Scan(&id)
	if err != nil {
		log.Printf("Error inserting patient: %s", err)
		return nil, err
	}

	patient := &entities.Patient{
		ID:                  p.ID,
		PhoneNumber:         p.PhoneNumber,
		Name:                p.Name,
		BirthDate:           p.BirthDate,
		GenderID:            p.GenderID,
		IdentityCardScanImg: p.IdentityCardScanImg,
	}

	return patient, nil
}
