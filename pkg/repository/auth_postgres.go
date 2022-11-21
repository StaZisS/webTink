package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
	listing "web"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user listing.User) (uuid.UUID, error) {
	var id uuid.UUID
	query := fmt.Sprintf("INSERT INTO %s (name, surname, username, email, grade, password, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Surname, user.Username, user.Email, user.Grade, user.Password, time.Now())
	if err := row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (listing.User, error) {
	var user listing.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password = $2", usersTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}

func (r *AuthPostgres) GetUserById(id uuid.UUID) (listing.User, error) {
	var user listing.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&user, query, id)
	return user, err
}
