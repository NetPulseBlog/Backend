package entity

import "github.com/google/uuid"

type ArticleCategory struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`

	AvatarUrl string `json:"avatar_url,omitempty"`
	CoverUrl  string `json:"cover_url,omitempty"`
}
