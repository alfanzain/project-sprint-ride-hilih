package repositories

import "github.com/alfanzain/project-sprint-halo-suster/app/entities"

type IUserRepository interface {
	FindByID(string) (*entities.User, error)
	FindByNIP(string) (*entities.User, error)
	DoesNIPExist(string) (bool, error)
	Store(*entities.UserStorePayload) (*entities.User, error)
	GetUsers(*entities.UserGetFilterParams) ([]*entities.User, error)
	Update(*entities.UserUpdatePayload) (*entities.UserUpdateResponse, error)
}
