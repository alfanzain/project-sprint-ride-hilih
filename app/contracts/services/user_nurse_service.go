package services

import "github.com/alfanzain/project-sprint-halo-suster/app/entities"

type IUserNurseService interface {
	Register(*entities.UserNurseRegisterPayload) (*entities.UserLoginResponse, error)
	Login(*entities.UserITLoginPayload) (*entities.UserLoginResponse, error)
	Update(*entities.UserUpdatePayload) (*entities.UserUpdateResponse, error)
	UpdatePassword(*entities.UserNurseGrantAccessPayload) (*entities.UserUpdatePasswordResponse, error)
	Delete(string) (bool, error)
}
