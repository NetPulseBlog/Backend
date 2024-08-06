package v1

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

//TODO: ADD AUTH MIDDLEWARE

// ListBookmarks A handler for getting a list of bookmarks.
func (h *Handler) ListBookmarks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Список закладок"))
}

// CreateBookmark Handler for creating a bookmark.
func (h *Handler) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resourceType := chi.URLParam(r, "type")
	w.Write([]byte("Создана закладка с ID: " + id + " & Тип: " + resourceType))
}

// DeleteBookmark Handler for deleting a bookmark.
func (h *Handler) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Удалена закладка с ID: " + id))
}
