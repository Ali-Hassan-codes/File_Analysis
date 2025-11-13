package services

import (
	"github.com/ali-hassan-Codes/file_analyzer_2/models"
	"github.com/ali-hassan-Codes/file_analyzer_2/repositories"
	"golang.org/x/crypto/bcrypt"
)

type SignupService struct {
	repo *repositories.UserRepository
}

func NewSignupService(repo *repositories.UserRepository) *SignupService {
	return &SignupService{repo: repo}
}

func (s *SignupService) Signup(user models.User) (models.User, error) {
	// Hash password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashed)
	return s.repo.CreateUser(user)
}
