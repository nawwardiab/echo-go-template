package repository

import (
	"echo-server/internal/model"
	"fmt"

	"github.com/jackc/pgx"
)

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) *UserRepo{
	return &UserRepo{db: db}
}

// GetByUserName uses db connection to query users table by username
func (r *UserRepo) GetByUserName(username string) (*model.User, error){
	u := &model.User{}
	query := `SELECT * FROM users WHERE username=$1`
	queryErr := r.db.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash)
	
	if queryErr != nil {
		return nil, fmt.Errorf("Get user by username: %w", queryErr)
	} else {
		return u, nil
	}
}

// CreateUser uses db connection to query users table and insert a new user
func (r *UserRepo) CreateUser(u *model.User) error {
	query := `INSERT Into users (username, email, password_hash) VALUES ($1, $2, $3)`
	_, execErr := r.db.Exec(query, u.Username, u.Email, u.PasswordHash)

	if execErr != nil {
		return fmt.Errorf("create user: %w", execErr)
	} else {
		return nil
	}
}
