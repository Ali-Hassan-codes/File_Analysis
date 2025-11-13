package services

import (
	"errors"

	"github.com/ali-hassan-Codes/file_analyzer_2/models"
	"github.com/ali-hassan-Codes/file_analyzer_2/repositories"
)

type LoginService struct {
	repo *repositories.UserRepository
}

func NewLoginService(repo *repositories.UserRepository) *LoginService {
	return &LoginService{repo: repo}
}

// Login method
func (s *LoginService) Login(email, password string) (models.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	// Compare password directly
	if user.Password != password {
		return models.User{}, errors.New("invalid credentials")
	}

	return user, nil
}
