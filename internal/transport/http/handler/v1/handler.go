package v1

import (
	"app/internal/config"
	"app/internal/service"
	"app/pkg/auth"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

type Handler struct {
	log      *slog.Logger
	cfg      *config.Config
	services *service.Services
}

func NewHandler(log *slog.Logger, cfg *config.Config, services *service.Services) *Handler {
	return &Handler{
		log:      log,
		cfg:      cfg,
		services: services,
	}
}

func (h *Handler) InitRouter() http.Handler {
	v1 := chi.NewRouter()

	authGuard := auth.CreateGuardMiddleware(h.log, h.services.Auth)

	v1.Route("/user", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign-up", h.UserSignUp)
			r.Post("/sign-in", h.UserSignIn)
			r.Post("/refresh-token", h.UserAuthTokenRefresh)
		})

		r.With(authGuard).Put("/settings", h.UserSettingsUpdate)
		r.With(authGuard).Put("/", h.UserEdit)

		r.With(authGuard).Post("/subscribe/{id}", h.UserSubscribe)
		r.With(authGuard).Post("/unsubscribe/{id}", h.UserUnsubscribe)

		r.With(authGuard).Get("/", h.UserProfile)
		r.Get("/{id}", h.UserProfileByID)

		r.Get("/sub-sites", h.UserSubSites)

		r.Route("/password", func(r chi.Router) {
			r.Post("/request_change", h.UserPasswordRequestChange)
			r.Post("/confirm_change", h.UserPasswordConfirmChange)
		})
	})

	v1.Route("/article", func(r chi.Router) {
		r.With(authGuard).Post("/", h.CreateArticle)

		r.Route("/list", func(r chi.Router) {
			r.Get("/{type}", h.ListArticles)
			r.With(authGuard).Get("/drafts", h.DraftListArticles)
		})

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.GetArticleByID)
			r.Delete("/", h.DeleteArticle)
			r.With(authGuard).Post("/subscribe", h.SubscribeToArticle)
			r.With(authGuard).Post("/unsubscribe", h.UnsubscribeFromArticle)
			r.With(authGuard).Put("/", h.EditArticle)
			r.With(authGuard).Post("/change-status", h.ChangeArticleStatus)
			r.Route("/comment", func(r chi.Router) {
				r.Get("/", h.GetCommentsForArticle)
				r.With(authGuard).Post("/", h.CreateCommentOnArticle)
				r.With(authGuard).Route("/{commentId}", func(r chi.Router) {
					r.Delete("/", h.DeleteCommentFromArticle)
					r.Put("/", h.EditCommentOnArticle)
				})
			})
		})
	})

	v1.With(authGuard).Route("/bookmark", func(r chi.Router) {
		r.Get("/list", h.ListBookmarks)
		r.Post("/{resource_id}/{resource_type}", h.CreateBookmark)
		r.Delete("/{resource_id}", h.DeleteBookmark)
	})

	return v1
}
