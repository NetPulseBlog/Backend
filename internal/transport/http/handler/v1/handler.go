package v1

import (
	"app/internal/service"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

type Handler struct {
	log      *slog.Logger
	services *service.Services
}

func NewHandler(log *slog.Logger, services *service.Services) *Handler {
	return &Handler{
		log:      log,
		services: services,
	}
}

func (h *Handler) InitRouter() http.Handler {
	v1 := chi.NewRouter()

	v1.Mount("/", h.initExampleRoutes())

	return v1
}
