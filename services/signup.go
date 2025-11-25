package services

import (
	"github.com/ali-hassan-Codes/file_analyzer_2/models"
	"github.com/ali-hassan-Codes/file_analyzer_2/repositories"
	"golang.org/x/crypto/bcrypt"
)

// Interface
type ISignupService interface {
	Signup(user models.User) (models.User, error)
}

// Concrete struct
type SignupService struct {
	repo repositories.UserRepositoryInterface
}

// Dependency struct for constructor
type SignupServiceDeps struct {
	Repo repositories.UserRepositoryInterface
}

// Constructor accepts dependency struct, returns interface
func NewSignupService(deps SignupServiceDeps) ISignupService {
	return &SignupService{
		repo: deps.Repo,
	}
}

// Business logic
func (s *SignupService) Signup(user models.User) (models.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashed)

	return s.repo.CreateUser(user)
}
