package services

import (
	"database/sql"
	"errors"
	"log"
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

func (s *UserNurseService) Register(p *entities.UserNurseRegisterPayload) (*entities.UserLoginResponse, error) {
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

	return &entities.UserLoginResponse{
		ID:          userIT.ID,
		NIP:         userIT.NIP,
		Name:        p.Name,
		AccessToken: accessToken,
	}, nil
}

func (s *UserNurseService) Login(p *entities.UserNurseLoginPayload) (*entities.UserLoginResponse, error) {
	log.Println("user nurse service login")
	decodedNIP, err := helpers.DecodeNIP(p.NIP)
	if err != nil {
		return nil, err
	}

	_, err = helpers.IsNIPNurseValid(decodedNIP.RoleID)
	if err != nil {
		return nil, errs.ErrUserNotFound
	}

	userNurse, err := s.userRepository.FindNurseByNIP(p.NIP)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ErrUserNotFound
		}

		return nil, err
	}

	isValidPassword := helpers.CheckPasswordHash(p.Password, *userNurse.Password)
	if !isValidPassword {
		return nil, errs.ErrInvalidPassword
	}

	paramsGenerateJWTLogin := helpers.ParamsGenerateJWT{
		SecretKey: os.Getenv("JWT_SECRET"),
		UserID:    userNurse.ID,
		RoleID:    userNurse.RoleID,
	}

	accessToken, err := helpers.GenerateJWT(&paramsGenerateJWTLogin)
	if err != nil {
		return nil, err
	}

	return &entities.UserLoginResponse{
		ID:          userNurse.ID,
		NIP:         userNurse.NIP,
		Name:        userNurse.Name,
		AccessToken: accessToken,
	}, nil
}

func (s *UserNurseService) Update(p *entities.UserUpdatePayload) (*entities.UserUpdateResponse, error) {
	user, err := s.userRepository.FindByID(p.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrUserNotFound
		}

		return nil, err
	}

	roleID, err := strconv.Atoi(user.RoleID)
	if err != nil {
		return nil, err
	}

	_, err = helpers.IsNIPNurseValid(roleID)
	if err != nil {
		return nil, errs.ErrUserNurseNotFound
	}

	decodedNIP, err := helpers.DecodeNIP(p.NIP)
	if err != nil {
		return nil, err
	}

	_, err = helpers.IsNIPNurseValid(decodedNIP.RoleID)
	if err != nil {
		return nil, err
	}

	updatedUser, err := s.userRepository.Update(p)
	if err != nil {
		return nil, err
	}

	return &entities.UserUpdateResponse{
		ID:   updatedUser.ID,
		NIP:  user.NIP,
		Name: updatedUser.Name,
	}, nil
}

func (s *UserNurseService) UpdatePassword(p *entities.UserNurseGrantAccessPayload) (*entities.UserUpdatePasswordResponse, error) {
	user, err := s.userRepository.FindByID(p.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ErrUserNotFound
		}

		return nil, err
	}

	roleID, err := strconv.Atoi(user.RoleID)
	if err != nil {
		return nil, err
	}

	_, err = helpers.IsNIPNurseValid(roleID)
	if err != nil {
		return nil, errs.ErrUserNurseNotFound
	}

	hashedPassword, err := helpers.HashPassword(p.Password)
	if err != nil {
		return nil, err
	}

	p.Password = hashedPassword

	updatedUser, err := s.userRepository.UpdatePassword(p)
	if err != nil {
		return nil, err
	}

	return &entities.UserUpdatePasswordResponse{
		ID: updatedUser.ID,
	}, nil
}

func (s *UserNurseService) Delete(userID string) (bool, error) {
	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errs.ErrUserNotFound
		}

		return false, err
	}

	roleID, err := strconv.Atoi(user.RoleID)
	if err != nil {
		return false, err
	}

	_, err = helpers.IsNIPNurseValid(roleID)
	if err != nil {
		return false, errs.ErrUserNurseNotFound
	}

	_, err = s.userRepository.Destroy(userID)
	if err != nil {
		return false, err
	}

	return true, nil
}
