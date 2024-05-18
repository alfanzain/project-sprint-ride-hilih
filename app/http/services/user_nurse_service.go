package services

import (
	"os"
	"strconv"

	repositoryContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/repositories"
	serviceContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/services"
	"github.com/alfanzain/project-sprint-halo-suster/app/entities"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/errs"
	"github.com/alfanzain/project-sprint-halo-suster/app/http/helpers"
)

type UserNurseService struct {
	userRepository repositoryContracts.IUserRepository
}

func NewUserNurseService(
	userRepository repositoryContracts.IUserRepository,
) serviceContracts.IUserNurseService {
	return &UserNurseService{
		userRepository: userRepository,
	}
}

func (s *UserNurseService) Register(p *entities.UserNurseRegisterPayload) (*entities.User, error) {
	nipExists, _ := s.userRepository.DoesNIPExist(p.NIP)

	if nipExists {
		return nil, errs.ErrNIPAlreadyRegistered
	}

	decodedNIP, err := helpers.DecodeNIP(p.NIP)
	if err != nil {
		return nil, err
	}

	_, err = helpers.IsNIPNurseValid(decodedNIP.RoleID)
	if err != nil {
		return nil, err
	}

	userIT, err := s.userRepository.Store(&entities.UserStorePayload{
		NIP:      p.NIP,
		Name:     p.Name,
		RoleID:   strconv.Itoa(decodedNIP.RoleID),
		GenderID: strconv.Itoa(decodedNIP.GenderID),
	})

	if err != nil {
		return nil, err
	}

	paramsGenerateJWTRegister := helpers.ParamsGenerateJWT{
		SecretKey: os.Getenv("JWT_SECRET"),
		UserID:    userIT.ID,
		RoleID:    userIT.RoleID,
	}

	accessToken, errAccessToken := helpers.GenerateJWT(&paramsGenerateJWTRegister)
	if errAccessToken != nil {
		return nil, errAccessToken
	}

	return &entities.User{
		ID:          userIT.ID,
		NIP:         userIT.NIP,
		Name:        p.Name,
		AccessToken: accessToken,
	}, nil
}

func (s *UserNurseService) Login(p *entities.UserITLoginPayload) (*entities.User, error) {
	userIT, err := s.userRepository.FindByNIP(p.NIP)
	if err != nil {
		return nil, err
	}
	if userIT == nil {
		return nil, errs.ErrUserITNotFound
	}

	isValidPassword := helpers.CheckPasswordHash(p.Password, userIT.Password)
	if !isValidPassword {
		return nil, errs.ErrInvalidPassword
	}

	paramsGenerateJWTLogin := helpers.ParamsGenerateJWT{
		SecretKey: os.Getenv("JWT_SECRET"),
		UserID:    userIT.ID,
		RoleID:    userIT.RoleID,
	}

	accessToken, err := helpers.GenerateJWT(&paramsGenerateJWTLogin)
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:          userIT.ID,
		NIP:         userIT.NIP,
		Name:        userIT.Name,
		AccessToken: accessToken,
	}, nil
}
