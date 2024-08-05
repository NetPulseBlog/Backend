package entity

import (
	"github.com/google/uuid"
	"time"
)

type BookmarkResourceType string

const (
	BTArticle BookmarkResourceType = "bt_article"
	BTComment BookmarkResourceType = "bt_comment"
)

type UserBookmark struct {
	UserId uuid.UUID `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`

	ResourceType BookmarkResourceType `json:"resource_type"`
	ResourceId   uuid.UUID            `json:"resource_id"`
}
