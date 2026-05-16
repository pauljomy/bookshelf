package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(app.enableCORS)
	mux.NotFound(app.notFoundResponse)

	mux.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)

	})

	return mux
}
