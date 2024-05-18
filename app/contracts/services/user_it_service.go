package services

import "github.com/alfanzain/project-sprint-halo-suster/app/entities"

type IUserITService interface {
	Register(*entities.UserITRegisterPayload) (*entities.UserLoginResponse, error)
	Login(*entities.UserITLoginPayload) (*entities.UserLoginResponse, error)
}
