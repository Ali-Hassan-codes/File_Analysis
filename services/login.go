package services

import (
	"errors"

	"github.com/ali-hassan-Codes/file_analyzer_2/models"
	"github.com/ali-hassan-Codes/file_analyzer_2/repositories"
	"golang.org/x/crypto/bcrypt"
)

// Interface
type ILoginService interface {
	Login(email, password string) (models.User, error)
}

// Concrete struct
type LoginService struct {
	repo repositories.UserRepositoryInterface
}

// Dependency struct
type LoginServiceDeps struct {
	Repo repositories.UserRepositoryInterface
}

// Constructor - returns interface
func NewLoginService(deps LoginServiceDeps) ILoginService {
	return &LoginService{
		repo: deps.Repo,
	}
}

// Business logic
func (s *LoginService) Login(email, password string) (models.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return models.User{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("invalid email or password")
	}

	return user, nil
}
