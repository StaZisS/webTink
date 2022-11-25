package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
	listing "web"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) Create(userId uuid.UUID, post listing.Post) (uuid.UUID, error) {
	var id uuid.UUID
	query := fmt.Sprintf("INSERT INTO %s (author_id, name, surname, role, education, additional, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", postsTable)
	row := r.db.QueryRow(query, userId, post.Name, post.Surname, post.Role, post.Education, post.Additional, time.Now(), time.Now())
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *PostPostgres) GetAll() ([]listing.PostSend, error) {
	var posts []listing.PostSend
	query := fmt.Sprintf("SELECT posts.id, posts.author_id, posts.name, posts.surname, posts.role, posts.education, posts.additional, posts.created_at, posts.updated_at FROM %s", postsTable)
	err := r.db.Select(&posts, query)
	return posts, err
}

func (r *PostPostgres) GetById(id uuid.UUID) (listing.PostSend, error) {
	var post listing.PostSend
	query := fmt.Sprintf("SELECT posts.id, posts.author_id, posts.name, posts.surname, posts.role, posts.education, posts.additional, posts.created_at, posts.updated_at FROM %s WHERE posts.id = $1", postsTable)
	err := r.db.Get(&post, query, id)
	return post, err
}

func (r *PostPostgres) Delete(idUser, idPost uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND author_id = $2", postsTable)
	_, err := r.db.Exec(query, idPost, idUser)
	return err
}

func (r *PostPostgres) Update(idUser, idPost uuid.UUID, input listing.UpdatePostInput) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1, surname = $2, role = $3, education = $4, additional = $5, updated_at = $7 WHERE id = $7 AND author_id = $8", postsTable)
	_, err := r.db.Exec(query, input.Name, input.Surname, input.Role, input.Education, input.Additional, time.Now(), idPost, idUser)
	return err
}
