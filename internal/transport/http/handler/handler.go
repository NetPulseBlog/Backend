package handler

import (
	"app/internal/service"
	v1 "app/internal/transport/http/handler/v1"
	"app/pkg/transport/http/middleware/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
)

type Handler struct {
	log      *slog.Logger
	services *service.Services
}

func NewTransportHandler(log *slog.Logger, services *service.Services) http.Handler {
	handler := Handler{
		log:      log,
		services: services,
	}

	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Logger)
	router.Use(logger.New(handler.log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// Api v1
	handlerV1 := v1.NewHandler(handler.log, handler.services)
	router.Route("/api", func(r chi.Router) {
		r.Mount("/v1", handlerV1.InitRouter())
	})

	return router
}
