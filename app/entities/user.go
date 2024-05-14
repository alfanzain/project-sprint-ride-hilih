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
