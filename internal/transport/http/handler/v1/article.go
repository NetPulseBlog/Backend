package v1

import (
	"app/internal/service/dto"
	"app/pkg/api/request"
	"app/pkg/api/response"
	"app/pkg/domain/entity"
	"app/pkg/infra/logger/sl"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

func (h *Handler) CreateDraftArticle(w http.ResponseWriter, r *http.Request) {
	// Логика создания черновика
	w.Write([]byte("Черновик статьи создан"))
}

func (h *Handler) PublishArticle(w http.ResponseWriter, r *http.Request) {
	// Логика публикации статьи
	w.Write([]byte("Статья опубликована"))
}

func (h *Handler) DraftListArticles(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) ListArticles(w http.ResponseWriter, r *http.Request) {
	// Логика фильтрации статей
	lType := chi.URLParam(r, "type")
	// Логика получения статьи по ID
	w.Write([]byte("Тип: " + lType))
}

func (h *Handler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// Логика получения статьи по ID
	w.Write([]byte("Статья с ID: " + id))
}

func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	const op = "http.v1.Article.DeleteArticle"

	id := chi.URLParam(r, "id")
	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	validate := validator.New()
	err := validate.Var(id, "required,uuid")
	if err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		log.Error("Invalid request", sl.Err(err))
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.ValidationError(validateErr))
		return
	}

	err = h.services.Article.Delete(id)
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.Error(response.ErrInternalServerError))
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (h *Handler) SubscribeToArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// Логика подписки на статью
	w.Write([]byte("Подписка на статью с ID: " + id))
}

func (h *Handler) UnsubscribeFromArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// Логика отписки от статьи
	w.Write([]byte("Отписка от статьи с ID: " + id))
}

func (h *Handler) EditArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// Логика редактирования статьи по ID
	w.Write([]byte("Статья с ID: " + id + " отредактирована"))
}

func (h *Handler) PublishArticleFromDraft(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// Логика публикации статьи из черновика по ID
	w.Write([]byte("Статья с ID: " + id + " опубликована из черновика"))
}

func (h *Handler) GetCommentsForArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Комментарии к статье с ID: " + id))
}

func (h *Handler) CreateCommentOnArticle(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "id")
	w.Write([]byte("Комментарий создан для статьи с ID: " + articleId))
}

func (h *Handler) DeleteCommentFromArticle(w http.ResponseWriter, r *http.Request) {
	commentId := chi.URLParam(r, "commentId")
	w.Write([]byte("Комментарий с ID: " + commentId + " удален"))
}

func (h *Handler) EditCommentOnArticle(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "id")
	commentId := chi.URLParam(r, "commentId")
	w.Write([]byte("Комментарий с ID: " + commentId + " отредактирован для статьи с ID: " + articleId))
}
