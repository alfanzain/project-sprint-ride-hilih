package services

import "github.com/alfanzain/project-sprint-halo-suster/app/entities"

type IUserNurseService interface {
	Register(*entities.UserNurseRegisterPayload) (*entities.User, error)
	Login(*entities.UserITLoginPayload) (*entities.User, error)
}
