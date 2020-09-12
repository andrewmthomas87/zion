package auth

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// NewService creates a new authentication service.
func NewService(dbpool *pgxpool.Pool) *Service {
	return &Service{dbpool: dbpool}
}

// Service implements authentication.
type Service struct {
	dbpool *pgxpool.Pool
}

// SignUp registers a new user.
func (s *Service) SignUp(ctx context.Context, username, password string) (*User, error) {
	sql := "INSERT INTO users (username, password) VALUES ($1, crypt($2, gen_salt('bf')))"
	if _, err := s.dbpool.Exec(ctx, sql, username, password); err != nil {
		return nil, err
	}

	sql = "SELECT id, username FROM users WHERE username=$1"
	row := s.dbpool.QueryRow(ctx, sql, username)
	var user *User
	if err := row.Scan(&user.ID, &user.Username); err != nil {
		return nil, err
	}
	return user, nil
}

// SignIn authenticates a user by username and password and returns the User if successful.
func (s *Service) SignIn(ctx context.Context, username, password string) (*User, error) {
	sql := "SELECT id, username FROM users WHERE username=$1 && password=crypt($2, password)"
	row := s.dbpool.QueryRow(ctx, sql, username, password)
	var user *User
	if err := row.Scan(&user.ID, &user.Username); err != nil {
		return nil, err
	}
	return user, nil
}

// User models a user.
type User struct {
	ID       int64
	Username string
}
