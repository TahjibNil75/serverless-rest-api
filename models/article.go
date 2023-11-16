package models

import "time"

type Article struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
