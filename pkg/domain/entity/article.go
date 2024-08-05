package entity

import (
	"github.com/google/uuid"
	"time"
)

type ArticleStatus string

const (
	ArticlePublishedStatus ArticleStatus = "published"
	ArticleDraftStatus     ArticleStatus = "draft"
)

type Article struct {
	Id       uuid.UUID `json:"id"`
	AuthorId uuid.UUID `json:"author_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Status ArticleStatus `json:"status"`

	Title   string    `json:"title"`
	TopicId uuid.UUID `json:"topic_id,omitempty"`

	ContentBlocks string `json:"content_blocks"`
	CoverUrl      string `json:"cover_url"`
	SubTitle      string `json:"sub_title"`

	ViewsCount    int `json:"views_count"`
	CommentsCount int `json:"comments_count"`
	BookmarkCount int `json:"bookmark_count"`
}
