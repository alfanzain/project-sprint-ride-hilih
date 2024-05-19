package entities

import "time"

type (
	Patient struct {
		ID                  int       `json:"identityNumber"`
		PhoneNumber         string    `json:"phoneNumber"`
		Name                string    `json:"name"`
		BirthDate           time.Time `json:"birthDate"`
		GenderID            int       `json:"gender"`
		IdentityCardScanImg string    `json:"-"`
		CreatedAt           time.Time `json:"createdAt"`
	}

	PatientRegisterRequest struct {
		ID                  int    `json:"identityNumber" validate:"required"`
		PhoneNumber         string `json:"phoneNumber" validate:"required,min=10,max=15"`
		Name                string `json:"name" validate:"required,min=3,max=30"`
		BirthDate           string `json:"birthDate" validate:"required"`
		Gender              string `json:"gender"  validate:"required"`
		IdentityCardScanImg string `json:"-"`
	}

	PatientStorePayload struct {
		ID                  int
		PhoneNumber         string
		Name                string
		BirthDate           time.Time
		GenderID            int
		IdentityCardScanImg string
	}
)
