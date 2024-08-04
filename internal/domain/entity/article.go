package entity

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	Id       uuid.UUID `json:"id"`
	AuthorId uuid.UUID `json:"authorId"`

	IsPublished bool `json:"isPublished"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Title    string          `json:"title"`
	Category ArticleCategory `json:"category,omitempty"`
	Content  string          `json:"content"` // todo think about blocks

	ViewsCount    int `json:"viewsCount"`
	CommentsCount int `json:"commentsCount"`
	BookmarkCount int `json:"bookmarkCount"`
}
