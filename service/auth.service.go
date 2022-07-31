package service

import (
	"log"

	"github.com/mashingan/smapping"
	"gitlab.com/mr687/privy-be-test-go/dto"
	"gitlab.com/mr687/privy-be-test-go/entity"
	"gitlab.com/mr687/privy-be-test-go/helper"
	"gitlab.com/mr687/privy-be-test-go/repository"
)

// Define contract this service can do
type AuthService interface {
	ValidateLogin(loginData dto.LoginDTO) (*entity.User, error)
	CreateUser(user dto.RegisterDTO) (*entity.User, error)
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (service *authService) ValidateLogin(loginData dto.LoginDTO) (*entity.User, error) {
	user, err := service.userRepo.FindByEmail(loginData.Email)
	if err != nil {
		return nil, err
	}
	if !helper.ValidateHash(loginData.Password, user.Password) {
		return nil, nil
	}
	return user, nil
}

func (service *authService) CreateUser(user dto.RegisterDTO) (*entity.User, error) {
	newUser := &entity.User{}

	// Mapping dto to entity
	err := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res, err := service.userRepo.InsertUser(newUser)
	if err != nil {
		log.Fatalf("Failed insert %v", err)
	}
	return res, nil
}

func (service *authService) IsDuplicateEmail(email string) bool {
	existsUser, err := service.userRepo.FindByEmail(email)
	if err != nil {
		return false
	}
	return existsUser != nil
}
