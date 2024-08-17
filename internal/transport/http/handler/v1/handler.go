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

		r.Get("/{id}", h.UserProfileByID)

		r.Get("/sub-sites", h.UserSubSites)
		r.Route("/password", func(r chi.Router) {
			r.Post("/request_change", h.UserPasswordRequestChange)
			r.Post("/confirm_change", h.UserPasswordConfirmChange)
		})
	})

	v1.Route("/article", func(r chi.Router) {
		r.With(authGuard).Post("/draft", h.CreateDraftArticle) // POST /article/draft
		r.With(authGuard).Post("/publish", h.PublishArticle)   // POST /article/publish

		r.Route("/list", func(r chi.Router) {
			r.Get("/", h.ListArticles)                     // GET /article/list?filter=popular, fresh or my
			r.With(authGuard).Get("/my", h.ListMyArticles) // GET /article/list/my
		})

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.GetArticleByID)                                     // GET /article/{id}
			r.With(authGuard).Post("/subscribe", h.SubscribeToArticle)       // POST /article/{id}/subscribe
			r.With(authGuard).Post("/unsubscribe", h.UnsubscribeFromArticle) // POST /article/{id}/unsubscribe
			r.With(authGuard).Put("/", h.EditArticle)                        // PUT /article/{id}
			r.With(authGuard).Post("/publish", h.PublishArticle)             // POST /article/{id}/publish
			r.Route("/comment", func(r chi.Router) {
				r.Get("/", h.GetCommentsForArticle)                   // GET /article/{id}/comment
				r.With(authGuard).Post("/", h.CreateCommentOnArticle) // POST /article/{id}/comment
				r.With(authGuard).Route("/{commentId}", func(r chi.Router) {
					r.Delete("/", h.DeleteCommentFromArticle) // DELETE /article/{id}/comment/{commentId}
					r.Put("/", h.EditCommentOnArticle)        // PUT /article/{id}/comment/{commentId}
				})
			})
		})
	})

	v1.With(authGuard).Route("/bookmark", func(r chi.Router) {
		r.Get("/list", h.ListBookmarks)                            // GET /bookmark/list
		r.Post("/{resource_id}/{resource_type}", h.CreateBookmark) // POST /bookmark/{resource_id}/{resource_type} // for found type see domain/user_bookmark.go
		r.Delete("/{resource_id}", h.DeleteBookmark)               // DELETE /bookmark/{resource_id}
	})

	return v1
}
