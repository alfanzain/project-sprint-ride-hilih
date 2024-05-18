package services

import (
	repositoryContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/repositories"
	serviceContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/services"
	"github.com/alfanzain/project-sprint-halo-suster/app/entities"
)

type UserService struct {
	userRepository repositoryContracts.IUserRepository
}

func NewUserService(
	userRepository repositoryContracts.IUserRepository,
) serviceContracts.IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetUsers(p *entities.UserGetFilterParams) ([]*entities.User, error) {
	users, err := s.userRepository.GetUsers(p)

	if err != nil {
		return nil, err
	}

	return users, nil
}
