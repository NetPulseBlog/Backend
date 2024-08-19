package entity

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type ArticleStatus string

const (
	ArticlePublishedStatus ArticleStatus = "published"
	ArticleDraftStatus     ArticleStatus = "draft"
)

const (
	ArticlePopularListType = "a_popular_lt"
	ArticleFreshListType   = "a_fresh_lt"
	ArticleMyListType      = "a_my_lt"
)

type Article struct {
	Id       uuid.UUID `json:"id"`
	AuthorId uuid.UUID `json:"author_id"`

	Author User `json:"author,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Status ArticleStatus `json:"status"`

	Title      string    `json:"title"`
	SubsSiteId uuid.UUID `json:"subs_site_id,omitempty"`

	ContentBlocks string `json:"content_blocks"`
	CoverUrl      string `json:"cover_url"`
	SubTitle      string `json:"sub_title"`

	ViewsCount    int `json:"views_count"`
	CommentsCount int `json:"comments_count"`
	BookmarkCount int `json:"bookmark_count"`
}

const ArticleStatusValidationField = "article_status"

func ArticleStatusValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	switch value {
	case
		string(ArticlePublishedStatus),
		string(ArticleDraftStatus):
		return true
	}

	return true
}
