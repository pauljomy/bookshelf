package main

import "net/http"

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {

	env := map[string]any{
		"error": message,
	}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logger.Error("Server error", "method", r.Method, "uri", r.URL.RequestURI())
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	app.logger.Warn("Not found", "method", r.Method, "uri", r.URL.RequestURI())
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}
