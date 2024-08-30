package dto

import (
	"github.com/google/uuid"
	"net/http"
)

type CreateArticleRequestDTO struct {
	SubsSiteId  uuid.UUID `json:"subsSiteId,omitempty"`
	Status      string    `json:"status" validate:"required,article_status"`
	Title       string    `json:"title" validate:"required"`
	Content     string    `json:"content" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Cover       http.File `json:"cover"`
}
