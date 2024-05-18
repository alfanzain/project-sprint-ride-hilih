package repositories

import "github.com/alfanzain/project-sprint-halo-suster/app/entities"

type IUserRepository interface {
	FindByNIP(string) (*entities.User, error)
	DoesNIPExist(string) (bool, error)
	Store(*entities.UserStorePayload) (*entities.User, error)
	GetUsers(*entities.UserGetFilterParams) ([]*entities.User, error)
}
