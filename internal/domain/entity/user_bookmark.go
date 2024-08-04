package entity

import (
	"github.com/google/uuid"
	"time"
)

const (
	BookmarkTypeArticle = iota + 1
	BookmarkTypeComment = iota + 1
)

type UserBookmark struct {
	UserId uuid.UUID `json:"userId"`

	CreatedAt time.Time `json:"createdAt"`

	ResourceType int       `json:"resourceType"`
	ResourceId   uuid.UUID `json:"resourceId"`
}
