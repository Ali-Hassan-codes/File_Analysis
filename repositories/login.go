package repositories

import (
	"database/sql"

	"github.com/ali-hassan-Codes/file_analyzer_2/models"
)

type LoginRepository struct {
	DB *sql.DB
}

func NewLoginRepository(db *sql.DB) *LoginRepository {
	return &LoginRepository{DB: db}
}

type LoginRepoInterface interface {
	GetByEmail(email string) (*models.User, error)
}

// Get user by email (for login)
func (r *LoginRepository) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE email=?"
	row := r.DB.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
