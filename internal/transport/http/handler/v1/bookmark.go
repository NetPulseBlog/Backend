package v1

import (
	"app/pkg/api/response"
	"app/pkg/domain/entity"
	"app/pkg/infra/logger/sl"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

// ListBookmarks A handler for getting a list of bookmarks.
func (h *Handler) ListBookmarks(w http.ResponseWriter, r *http.Request) {
	const op = "http.v1.ListBookmarks"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	resourceType := r.URL.Query().Get("resource_type")
	if resourceType == "" {
		resourceType = string(entity.BTArticle)
	}

	list, err := h.services.Bookmark.GetList(entity.BookmarkResourceType(resourceType))
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.JSON(w, r, response.Error(response.ErrInternalServerError))
		return
	}

	render.JSON(w, r, list)
}

// CreateBookmark Handler for creating a bookmark.
func (h *Handler) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	const op = "http.v1.CreateBookmark"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "resource_id")
	resourceType := chi.URLParam(r, "resource_type")
	authId, _ := r.Cookie(entity.AuthIdFieldName) // TODO: get user id from context

	err := h.services.Bookmark.Create(authId.Value, id, entity.BookmarkResourceType(resourceType))
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.JSON(w, r, response.Error(response.ErrInternalServerError))
		return
	}

	render.JSON(w, r, response.OK())
}

// DeleteBookmark Handler for deleting a bookmark.
func (h *Handler) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	const op = "http.v1.DeleteBookmark"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	_ = log
	id := chi.URLParam(r, "resource_id")

	err := h.services.Bookmark.Delete(id)
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.JSON(w, r, response.Error(response.ErrUnknownId))
		return
	}

	render.JSON(w, r, response.OK())
}
