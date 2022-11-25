package repository

import (
	"github.com/jmoiron/sqlx"
	listing "web"
)

type EmailPostgres struct {
	db *sqlx.DB
}

func NewEmailPostgres(db *sqlx.DB) *EmailPostgres {
	return &EmailPostgres{db: db}
}

func (r *EmailPostgres) SendEmail(email listing.Email) error {
	return nil
}
