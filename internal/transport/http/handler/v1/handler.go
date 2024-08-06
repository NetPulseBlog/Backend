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

	v1.Route("/user", func(r chi.Router) {
		//TODO: ADD AUTH MIDDLEWARE WHERE IT NEED .WITH(AUTH_MIDDLEWARE)

		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", h.UserAuthRegister)
			r.Post("/login", h.UserAuthLogin)
			r.Route("/token", func(r chi.Router) {
				r.Post("/refresh", h.UserAuthTokenRefresh)
			})
		})
		r.Route("/settings", func(r chi.Router) {
			r.Get("/news-line", h.UserSettingsNewsline)
		})
		r.Put("/", h.UserEdit)
		r.Get("/", h.UserProfile)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.UserProfileByID)
			r.Post("/subscribe", h.UserSubscribe)
			r.Post("/unsubscribe", h.UserUnsubscribe)
		})
		r.Get("/sub-sites", h.UserSubSites)
		r.Route("/password", func(r chi.Router) {
			r.Post("/request_change", h.UserPasswordRequestChange)
			r.Post("/confirm_change", h.UserPasswordConfirmChange)
		})
	})

	v1.Route("/article", func(r chi.Router) {
		//TODO: ADD AUTH MIDDLEWARE WHERE IT NEED .WITH(AUTH_MIDDLEWARE)

		r.Post("/draft", h.CreateDraftArticle) // POST /article/draft
		r.Post("/publish", h.PublishArticle)   // POST /article/publish

		r.Route("/list", func(r chi.Router) {
			r.Get("/", h.ListArticles)     // GET /article/list?filter=popular, fresh or my
			r.Get("/my", h.ListMyArticles) // GET /article/list/my
		})

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.GetArticleByID)                     // GET /article/{id}
			r.Post("/subscribe", h.SubscribeToArticle)       // POST /article/{id}/subscribe
			r.Post("/unsubscribe", h.UnsubscribeFromArticle) // POST /article/{id}/unsubscribe
			r.Put("/", h.EditArticle)                        // PUT /article/{id}
			r.Post("/publish", h.PublishArticle)             // POST /article/{id}/publish
			r.Route("/comment", func(r chi.Router) {
				r.Get("/", h.GetCommentsForArticle)   // GET /article/{id}/comment
				r.Post("/", h.CreateCommentOnArticle) // POST /article/{id}/comment
				r.Route("/{commentId}", func(r chi.Router) {
					r.Delete("/", h.DeleteCommentFromArticle) // DELETE /article/{id}/comment/{commentId}
					r.Put("/", h.EditCommentOnArticle)        // PUT /article/{id}/comment/{commentId}
				})
			})
		})
	})

	v1.Route("/bookmark", func(r chi.Router) {
		//TODO: ADD AUTH MIDDLEWARE
		r.Get("/list", h.ListBookmarks)          // GET /bookmark/list
		r.Post("/{id}/{type}", h.CreateBookmark) // POST /bookmark/{id}
		r.Delete("/{id}", h.DeleteBookmark)      // DELETE /bookmark/{id}
	})

	return v1
}
