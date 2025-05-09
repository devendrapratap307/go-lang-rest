package users

import (
	"database/sql"
	"fmt"

	"github.com/restapi-go/types"
)

type Store struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	user := new(types.User)

	if err := s.db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("query database error")
	}
	return user, nil
}
func (s *Store) CreateUser(user *types.User) error {
	_, err := s.db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("user insert database error")
	}
	return nil
}
