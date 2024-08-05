package entity

import (
	"github.com/google/uuid"
	"time"
)

type ArticleComment struct {
	Id             uuid.UUID `json:"id"`
	ArticleId      uuid.UUID `json:"article_id"`
	ReplyCommentId uuid.UUID `json:"reply_comment_id"`

	AuthorId uuid.UUID `json:"author_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Content string `json:"content"`

	IsEdited bool `json:"is_edited"`
}
