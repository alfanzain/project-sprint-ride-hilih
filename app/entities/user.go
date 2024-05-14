package entities

type (
	User struct {
		ID          string `json:"userId"`
		Name        string `json:"name"`
		NIP         string `json:"nip"`
		Password    string `json:"-"`
		RoleID      int    `json:"-"`
		GenderID    int    `json:"-"`
		AccessToken string `json:"accessToken"`
	}

	UserITRegisterRequest struct {
		NIP      int    `json:"nip" validate:"required"`
		Name     string `json:"name" validate:"required,min=5,max=50"`
		Password string `json:"password" validate:"required,min=5,max=15"`
	}

	UserITLoginRequest struct {
		NIP      int    `json:"nip" validate:"required,min=13,max=13"`
		Password string `json:"password" validate:"required,min=5,max=15"`
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

	UserITStorePayload struct {
		ID       string
		Name     string
		NIP      string
		RoleID   int
		GenderID int
		Password string
	}
)
