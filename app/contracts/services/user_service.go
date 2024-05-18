package services

import "github.com/alfanzain/project-sprint-halo-suster/app/entities"

type IUserService interface {
	GetUsers(*entities.UserGetFilterParams) ([]*entities.User, error)
}
