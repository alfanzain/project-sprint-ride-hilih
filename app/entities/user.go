package entities

import "time"

type (
	User struct {
		ID        string    `json:"userId"`
		Name      string    `json:"name"`
		NIP       int       `json:"nip"`
		Password  string    `json:"-"`
		RoleID    string    `json:"-"`
		GenderID  string    `json:"-"`
		CreatedAt time.Time `json:"createdAt"`
	}

	UserGetRequest struct {
		ID     string `param:"id"`
		Limit  string `param:"limit" validate:"omitempty,number"`
		Offset string `param:"offset" validate:"omitempty,number"`
		Name   string `param:"name"`
	}

	UserGetFilterParams struct {
		ID        string
		Limit     int
		Offset    int
		Name      string
		NIP       string
		Role      string
		CreatedAt string
	}

	UserUpdateRequest struct {
		NIP  int    `json:"nip" validate:"required"`
		Name string `json:"name" validate:"required,min=5,max=50"`
	}

	UserITRegisterRequest struct {
		NIP      int    `json:"nip" validate:"required"`
		Name     string `json:"name" validate:"required,min=5,max=50"`
		Password string `json:"password" validate:"required,min=5,max=15"`
	}

	UserITLoginRequest struct {
		NIP      int    `json:"nip" validate:"required"`
		Password string `json:"password" validate:"required,min=5,max=15"`
	}

	UserNurseRegisterRequest struct {
		NIP                 int    `json:"nip" validate:"required"`
		Name                string `json:"name" validate:"required,min=5,max=50"`
		IdentityCardScanImg string `json:"identityCardScanImg" validate:"required"` // should be url
	}

	UserNurseGrantAccessRequest struct {
		Password string `json:"password" validate:"required"`
	}

	UserITRegisterPayload struct {
		NIP      string
		Name     string
		Password string
	}

	UserITLoginPayload struct {
		NIP      string
		Password string
	}

	UserStorePayload struct {
		ID       string
		Name     string
		NIP      string
		RoleID   string
		GenderID string
		Password *string
	}

	UserNurseRegisterPayload struct {
		NIP                 string
		Name                string
		IdentityCardScanImg string
	}

	UserUpdatePayload struct {
		ID   string
		NIP  string
		Name string
	}

	UserLoginResponse struct {
		ID          string `json:"userId"`
		Name        string `json:"name"`
		NIP         int    `json:"nip"`
		AccessToken string `json:"accessToken"`
	}

	UserUpdateResponse struct {
		ID   string `json:"userId"`
		Name string `json:"name"`
		NIP  int    `json:"nip"`
	}
)
