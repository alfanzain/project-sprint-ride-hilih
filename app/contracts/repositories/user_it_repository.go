package repositories

import "github.com/alfanzain/project-sprint-halo-suster/app/entities"

// TODO: this contract should be IUserRepository
type IUserITRepository interface {
	FindByNIP(string) (*entities.User, error)
	DoesNIPExist(string) (bool, error)
	Store(*entities.UserITStorePayload) (*entities.User, error)
}
