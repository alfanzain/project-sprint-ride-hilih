package repositories

import "github.com/alfanzain/project-sprint-halo-suster/app/entities"

type IUserRepository interface {
	FindByID(string) (*entities.User, error)
	FindByNIP(string) (*entities.User, error)
	FindITByNIP(string) (*entities.User, error)
	FindNurseByNIP(string) (*entities.User, error)
	DoesNIPExist(string) (bool, error)
	Store(*entities.UserStorePayload) (*entities.User, error)
	GetUsers(*entities.UserGetFilterParams) ([]*entities.User, error)
	Update(*entities.UserUpdatePayload) (*entities.UserUpdateResponse, error)
	UpdatePassword(*entities.UserNurseGrantAccessPayload) (*entities.UserUpdatePasswordResponse, error)
	Destroy(string) (bool, error)
}
