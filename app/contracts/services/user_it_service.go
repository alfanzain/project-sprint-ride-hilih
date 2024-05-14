package services

import "github.com/alfanzain/project-sprint-halo-suster/app/entities"

type IUserITService interface {
	Register(*entities.UserITRegisterPayload) (*entities.User, error)
	Login(*entities.UserITLoginPayload) (*entities.User, error)
}
