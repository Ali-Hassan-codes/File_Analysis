package repositories

import (
	"database/sql"
	"errors"

	"github.com/ali-hassan-Codes/file_analyzer_2/models"
)

type UserRepositoryInterface interface {
	CreateUser(user models.User) (models.User, error)
	GetByEmail(email string) (models.User, error)
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{DB: db}
}


func (repo *UserRepository) CreateUser(user models.User) (models.User, error) {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err := repo.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	query := "SELECT id, name, email, password FROM users WHERE email=?"
	row := repo.DB.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	return user, nil
}
