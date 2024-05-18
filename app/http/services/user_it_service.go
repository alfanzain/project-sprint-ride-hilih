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

type UserITService struct {
	userITRepository repositoryContracts.IUserRepository
}

func NewUserITService(
	userITRepository repositoryContracts.IUserRepository,
) serviceContracts.IUserITService {
	return &UserITService{
		userITRepository: userITRepository,
	}
}

func (s *UserITService) Register(p *entities.UserITRegisterPayload) (*entities.UserLoginResponse, error) {
	nipExists, _ := s.userITRepository.DoesNIPExist(p.NIP)

	if nipExists {
		return nil, errs.ErrNIPAlreadyRegistered
	}

	hashedPassword, err := helpers.HashPassword(p.Password)
	if err != nil {
		return nil, err
	}

	decodedNIP, err := helpers.DecodeNIP(p.NIP)
	if err != nil {
		return nil, err
	}

	_, err = helpers.IsNIPITValid(decodedNIP.RoleID)
	if err != nil {
		return nil, err
	}

	userIT, err := s.userITRepository.Store(&entities.UserStorePayload{
		NIP:      p.NIP,
		Name:     p.Name,
		RoleID:   strconv.Itoa(decodedNIP.RoleID),
		GenderID: strconv.Itoa(decodedNIP.GenderID),
		Password: &hashedPassword,
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

	return &entities.UserLoginResponse{
		ID:          userIT.ID,
		NIP:         userIT.NIP,
		Name:        p.Name,
		AccessToken: accessToken,
	}, nil
}

func (s *UserITService) Login(p *entities.UserITLoginPayload) (*entities.UserLoginResponse, error) {
	userIT, err := s.userITRepository.FindByNIP(p.NIP)
	if err != nil {
		return nil, err
	}
	if userIT == nil {
		return nil, errs.ErrUserITNotFound
	}

	isValidPassword := helpers.CheckPasswordHash(p.Password, *userIT.Password)
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

	return &entities.UserLoginResponse{
		ID:          userIT.ID,
		NIP:         userIT.NIP,
		Name:        userIT.Name,
		AccessToken: accessToken,
	}, nil
}
