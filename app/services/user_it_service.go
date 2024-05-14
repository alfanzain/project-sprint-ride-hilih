package services

import (
	"errors"
	"os"

	repositoryContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/repositories"
	serviceContracts "github.com/alfanzain/project-sprint-halo-suster/app/contracts/services"
	"github.com/alfanzain/project-sprint-halo-suster/app/entities"
	"github.com/alfanzain/project-sprint-halo-suster/app/helpers"
)

type UserITService struct {
	userITRepository repositoryContracts.IUserITRepository
}

func NewUserITService(
	userITRepository repositoryContracts.IUserITRepository,
) serviceContracts.IUserITService {
	return &UserITService{
		userITRepository: userITRepository,
	}
}

var ErrNIPAlreadyRegistered = errors.New("nip already registered")
var ErrUserITNotFound = errors.New("user IT not found")
var ErrInvalidPassword = errors.New("invalid password")

func (s *UserITService) Register(p *entities.UserITRegisterPayload) (*entities.User, error) {
	nipExists, _ := s.userITRepository.DoesNIPExist(p.NIP)

	if nipExists {
		return nil, ErrNIPAlreadyRegistered
	}

	hashedPassword, _ := helpers.HashPassword(p.Password)
	decodedNIP, _ := helpers.DecodeNIP(p.Password)

	// if decodedNIP.RoleID ==

	userIT, err := s.userITRepository.Store(&entities.UserITStorePayload{
		NIP:      p.NIP,
		Name:     p.Name,
		RoleID:   decodedNIP.RoleID,
		GenderID: decodedNIP.GenderID,
		Password: hashedPassword,
	})

	paramsGenerateJWTRegister := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 480,
		SecretKey:       os.Getenv("JWT_SECRET"),
		UserId:          userIT.ID,
	}

	accessToken, errAccessToken := helpers.GenerateJWT(&paramsGenerateJWTRegister)

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:          userIT.ID,
		NIP:         p.NIP,
		Name:        p.Name,
		AccessToken: accessToken,
	}, nil
}

func (s *UserITService) Login(p *entities.UserITLoginPayload) (*entities.User, error) {
	userIT, err := s.userITRepository.FindByNIP(p.NIP)
	if err != nil {
		return nil, err
	}
	if userIT == nil {
		return nil, ErrUserITNotFound
	}

	isValidPassword := helpers.CheckPasswordHash(p.Password, userIT.Password)
	if !isValidPassword {
		return nil, ErrInvalidPassword
	}

	paramsGenerateJWTLogin := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 480,
		SecretKey:       os.Getenv("JWT_SECRET"),
		UserId:          userIT.ID,
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
