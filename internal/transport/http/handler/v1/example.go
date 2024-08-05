package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) initExampleRoutes() http.Handler {
	exampleRouter := chi.NewRouter()

	exampleRouter.Route("/example", func(r chi.Router) {
		r.Get("/", h.exampleHandler())
	})

	return exampleRouter
}

func (h *Handler) exampleHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		h.log.Info("Got URL")

		h.services.Example.Create("example name")

		render.HTML(writer, request, "<h1>Hello world</h1>")
	}
}
