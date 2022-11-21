package listing

import "github.com/google/uuid"

type Post struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	AuthorID uuid.UUID `json:"author_id"`
	Photo    string    `json:"photo"`
}
