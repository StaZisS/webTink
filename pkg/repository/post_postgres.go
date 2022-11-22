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
	query := fmt.Sprintf("INSERT INTO %s (title, content, author_id, photo, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", postsTable)
	row := r.db.QueryRow(query, post.Title, post.Content, userId, post.Photo, time.Now(), time.Now())
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *PostPostgres) GetAll() ([]listing.Post, error) {
	var posts []listing.Post
	query := fmt.Sprintf("SELECT posts.id, posts.title, posts.content, posts.author_id, posts.photo, posts.created_at, posts.updated_at FROM %s", postsTable)
	err := r.db.Select(&posts, query)
	return posts, err
}

func (r *PostPostgres) GetById(id uuid.UUID) (listing.Post, error) {
	var post listing.Post
	query := fmt.Sprintf("SELECT posts.id, posts.title, posts.content, posts.author_id, posts.photo, posts.created_at, posts.updated_at FROM %s WHERE posts.id = $1", postsTable)
	err := r.db.Get(&post, query, id)
	return post, err
}
