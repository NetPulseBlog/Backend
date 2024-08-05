package entity

import (
	"github.com/google/uuid"
	"time"
)

const (
	BTArticle = "bt_article"
	BTComment = "bt_comment"
)

type UserBookmark struct {
	UserId uuid.UUID `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`

	ResourceType int       `json:"resource_type"`
	ResourceId   uuid.UUID `json:"resource_id"`
}
