package services

import (
	"errors"

	"github.com/ali-hassan-Codes/file_analyzer_2/models"
	"github.com/ali-hassan-Codes/file_analyzer_2/repositories"
	"golang.org/x/crypto/bcrypt"
)

type ILoginService interface {
	Login(email, password string) (models.User, error)
}

// Service
type LoginService struct {
	repo *repositories.UserRepository
}

func NewLoginService(repo *repositories.UserRepository) *LoginService {
	return &LoginService{repo: repo}
}

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
