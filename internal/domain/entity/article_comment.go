package entity

import (
	"github.com/google/uuid"
	"time"
)

type ArticleComment struct {
	Id             uuid.UUID `json:"id"`
	ArticleId      uuid.UUID `json:"articleId"`
	ReplyCommentId uuid.UUID `json:"replyCommentId"`

	AuthorId uuid.UUID `json:"authorId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Content string `json:"content"`

	IsEdited  bool `json:"isEdited"`
	IsRemoved bool `json:"isRemoved"`
}
