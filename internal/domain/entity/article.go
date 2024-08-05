package entity

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	Id       uuid.UUID `json:"id"`
	AuthorId uuid.UUID `json:"author_id"`

	IsPublished bool `json:"is_published"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Title    string          `json:"title"`
	Category ArticleCategory `json:"category,omitempty"`

	Content  string `json:"content"`
	CoverUrl string `json:"cover_url"`
	SubTitle string `json:"sub_title"`

	ViewsCount    int `json:"views_count"`
	CommentsCount int `json:"comments_count"`
	BookmarkCount int `json:"bookmark_count"`
}
